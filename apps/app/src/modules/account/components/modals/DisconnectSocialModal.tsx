import { useState } from "react";
import { Dialog, Button } from "@lh/ui";
import { supabase } from "@/lib/supabase";

interface DisconnectSocialModalProps {
  isOpen: boolean;
  onClose: () => void;
  provider: string; // e.g. 'google'
  onSuccess?: () => void;
}

export function DisconnectSocialModal({ isOpen, onClose, provider, onSuccess }: DisconnectSocialModalProps) {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [errorMsg, setErrorMsg] = useState("");

  const handleDisconnect = async () => {
    setIsSubmitting(true);
    setErrorMsg("");
    
    try {
      const { data: { session } } = await supabase.auth.getSession();
      
      // 1. Unlink from Supabase client-side FIRST
      const { data: { user } } = await supabase.auth.getUser();
      const identity = user?.identities?.find((i) => i.provider === provider);
      
      if (identity) {
        const { error } = await supabase.auth.unlinkIdentity(identity);
        if (error) throw error;
      }

      // 2. Call our backend to convert the source_provider to 'manual'
      const res = await fetch(`http://localhost:3000/api/users/me/providers/${provider}/disconnect`, {
        method: "POST",
        headers: {
          "Authorization": `Bearer ${session?.access_token}`
        }
      });
      
      if (!res.ok) {
        const data = await res.json();
        // Even if backend fails, Supabase identity is unlinked. We still throw to show error.
        throw new Error(data.error || "Failed to sync provider disconnection with server");
      }
      
      onSuccess?.();
      onClose();
    } catch (err: any) {
      setErrorMsg(err.message);
    } finally {
      setIsSubmitting(false);
    }
  };

  const providerName = provider.charAt(0).toUpperCase() + provider.slice(1);

  return (
    <Dialog 
      isOpen={isOpen} 
      onClose={onClose}
      title={`Disconnect ${providerName}?`}
      description={`Are you sure you want to unlink your ${providerName} account?`}
    >
      <div className="mt-4 space-y-4 text-sm text-white/70">
        <p>
          You will no longer be able to use the 1-click <strong>{providerName}</strong> sign-in method to access your account.
        </p>
        
        <div className="bg-amber-900/20 border border-amber-500/30 rounded-lg p-4 space-y-2">
          <p className="flex items-start gap-2">
            <span className="text-amber-500 mt-0.5">⚠</span>
            <span>
              <strong>Security Recommendation:</strong> We highly recommend keeping multiple login methods active. Reducing your available sign-in options increases the risk of losing access to your account if you lose access to a primary device or email.
            </span>
          </p>
        </div>

        <p className="text-xs text-white/50">
          Note: Disconnecting this provider does not remove the associated email from your Limitless Hoops profile. You can still request a one-time passcode to that email address to sign in.
        </p>
      </div>

      {errorMsg && (
        <div className="mt-6 bg-red-900/20 border border-red-500/50 rounded-lg p-3">
          <p className="text-sm text-red-400">{errorMsg}</p>
        </div>
      )}
      <div className="mt-8 flex justify-end gap-3">
        <Button variant="ghost" onClick={onClose} disabled={isSubmitting} className="text-white hover:bg-white/10">
          Cancel
        </Button>
        <Button variant="destructive" onClick={handleDisconnect} disabled={isSubmitting}>
          {isSubmitting ? "Disconnecting..." : "Disconnect"}
        </Button>
      </div>
    </Dialog>
  );
}
