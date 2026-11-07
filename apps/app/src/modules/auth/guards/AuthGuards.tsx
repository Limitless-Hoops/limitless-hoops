import type { ReactNode } from "react";
import { Navigate } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";

export function ProtectedRoute({ children }: { children: ReactNode }) {
  const { session, household, loading } = useAuth();
  
  if (loading) return <div className="min-h-screen bg-black flex items-center justify-center text-white">Loading...</div>;
  
  if (!session) {
    return <Navigate to="/login" replace />;
  }

  if (!household) {
    return <Navigate to="/auth/onboarding" replace />;
  }

  return <>{children}</>;
}

export function PublicRoute({ children }: { children: ReactNode }) {
  const { session, household, loading } = useAuth();
  
  if (loading) return <div className="min-h-screen bg-black flex items-center justify-center text-white">Loading...</div>;
  
  if (session) {
    if (household) {
      return <Navigate to="/dashboard" replace />;
    } else {
      return <Navigate to="/auth/onboarding" replace />;
    }
  }

  return <>{children}</>;
}

export function OnboardingRoute({ children }: { children: ReactNode }) {
  const { session, household, loading } = useAuth();
  
  if (loading) return <div className="min-h-screen bg-black flex items-center justify-center text-white">Loading...</div>;
  
  if (!session) {
    return <Navigate to="/login" replace />;
  }

  if (household) {
    return <Navigate to="/dashboard" replace />;
  }

  return <>{children}</>;
}
