import { useState } from "react";
import { Dialog } from "@lh/ui";
import VerifyCodeForm from "@/modules/auth/components/VerifyCodeForm";
import { supabase } from "@/lib/supabase";

interface VerifyContactModalProps {
  isOpen: boolean;
  onClose: () => void;
  contactId: string;
  contactValue: string;
  contactType: "emails" | "phones";
  onSuccess: () => void;
}

export function VerifyContactModal({ isOpen, onClose, contactId, contactValue, contactType, onSuccess }: VerifyContactModalProps) {
  const [isVerifying, setIsVerifying] = useState(false);
  const [isResending, setIsResending] = useState(false);
  const [errorMsg, setErrorMsg] = useState("");
  const [resendSuccess, setResendSuccess] = useState(false);

  const getAuthHeader = async () => {
    const { data: { session } } = await supabase.auth.getSession();
    return { Authorization: `Bearer ${session?.access_token}` };
  };

  const handleVerify = async (code: string) => {
    setIsVerifying(true);
    setErrorMsg("");
    setResendSuccess(false);

    try {
      const headers = await getAuthHeader();
      const res = await fetch(`http://localhost:3000/api/users/me/${contactType}/${contactId}/verify`, {
        method: "POST",
        headers: { ...headers, "Content-Type": "application/json" },
        body: JSON.stringify({ code }),
      });
      
      if (!res.ok) {
        const data = await res.json();
        throw new Error(data.error || "Verification failed");
      }
      
      onSuccess();
      onClose();
    } catch (err: any) {
      setErrorMsg(err.message);
    } finally {
      setIsVerifying(false);
    }
  };

  const handleResend = async () => {
    setIsResending(true);
    setErrorMsg("");
    setResendSuccess(false);

    try {
      const headers = await getAuthHeader();
      const res = await fetch(`http://localhost:3000/api/users/me/${contactType}/${contactId}/resend`, {
        method: "POST",
        headers,
      });
      
      if (!res.ok) {
        throw new Error("Failed to resend code");
      }
      
      setResendSuccess(true);
    } catch (err: any) {
      setErrorMsg(err.message);
    } finally {
      setIsResending(false);
    }
  };

  return (
    <Dialog isOpen={isOpen} onClose={onClose}>
      <div className="flex flex-col gap-4">
        {errorMsg && <p className="text-sm text-red-400 text-center">{errorMsg}</p>}
        {resendSuccess && <p className="text-sm text-green-400 text-center">✅ Code resent successfully.</p>}
        <VerifyCodeForm 
          identifier={contactValue} 
          onVerify={handleVerify} 
          onResend={handleResend}
          loading={isVerifying}
        />
        {isResending && <p className="text-xs text-white/50 text-center">Sending new code...</p>}
      </div>
    </Dialog>
  );
}
