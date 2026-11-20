import { createContext, useContext, useEffect, useState } from "react";
import type { ReactNode } from "react";
import { supabase } from "@/lib/supabase";
import type { Session } from "@supabase/supabase-js";

interface AuthContextType {
  session: Session | null;
  household: any | null;
  loading: boolean;
  fetchError: string | null;
  refreshHousehold: () => Promise<void>;
  signOut: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType>({
  session: null,
  household: null,
  loading: true,
  fetchError: null,
  refreshHousehold: async () => {},
  signOut: async () => {},
});

export function AuthProvider({ children }: { children: ReactNode }) {
  const [session, setSession] = useState<Session | null>(null);
  const [household, setHousehold] = useState<any | null>(null);
  const [loading, setLoading] = useState(true);
  // H10/M14: surface fetch errors instead of silent null
  const [fetchError, setFetchError] = useState<string | null>(null);

  const fetchHousehold = async (token: string) => {
    try {
      const res = await fetch("http://localhost:3000/api/households/me", {
        headers: {
          "Authorization": `Bearer ${token}`
        }
      });
      if (res.ok) {
        const data = await res.json();
        setHousehold(data ?? null);
        setFetchError(null);
      } else if (res.status === 404 || res.status === 204) {
        // User exists but has no household yet — this is normal for new signups
        setHousehold(null);
        setFetchError(null);
      } else {
        // Actual server error — surface it so guards don't redirect incorrectly
        setHousehold(null);
        setFetchError(`Server error (${res.status}). Please refresh.`);
      }
    } catch (err) {
      console.error("fetchHousehold network error:", err);
      setHousehold(null);
      setFetchError("Network error. Please check your connection and refresh.");
    }
  };

  // M13: toggle loading during refreshHousehold so guards wait correctly
  const refreshHousehold = async () => {
    setLoading(true);
    const { data: { session: freshSession } } = await supabase.auth.getSession();
    if (freshSession?.access_token) {
      await fetchHousehold(freshSession.access_token);
    } else {
      setHousehold(null);
    }
    setLoading(false);
  };

  useEffect(() => {
    const initializeAuth = async () => {
      setLoading(true);
      const { data: { session } } = await supabase.auth.getSession();
      setSession(session);
      
      if (session?.access_token) {
        await fetchHousehold(session.access_token);
      }
      setLoading(false);
    };

    initializeAuth();

    const { data: authListener } = supabase.auth.onAuthStateChange(
      async (_event, newSession) => {
        // H10: set loading=true while we resolve the household to prevent
        // guard flash where loading=false + household=null => redirect to onboarding
        setLoading(true);
        setSession(newSession);
        if (newSession?.access_token) {
          await fetchHousehold(newSession.access_token);
        } else {
          setHousehold(null);
          setFetchError(null);
        }
        setLoading(false);
      }
    );

    return () => {
      authListener.subscription.unsubscribe();
    };
  }, []);

  const signOut = async () => {
    await supabase.auth.signOut();
  };

  return (
    <AuthContext.Provider value={{ session, household, loading, fetchError, refreshHousehold, signOut }}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => useContext(AuthContext);
