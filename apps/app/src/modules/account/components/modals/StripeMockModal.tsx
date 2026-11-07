
import { Dialog, Button } from "@lh/ui";

interface StripeMockModalProps {
  isOpen: boolean;
  onClose: () => void;
  action: "upgrade" | "cancel";
}

export function StripeMockModal({ isOpen, onClose, action }: StripeMockModalProps) {
  const isUpgrade = action === "upgrade";
  
  return (
    <Dialog 
      isOpen={isOpen} 
      onClose={onClose}
      title={isUpgrade ? "Redirect to Stripe Checkout" : "Redirect to Stripe Portal"}
      description={isUpgrade 
        ? "This will redirect you to the Stripe Customer Portal or Checkout to upgrade your plan." 
        : "This will redirect you to the Stripe Customer Portal where you can manage or cancel your active subscription."}
    >
      <div className="mt-8 flex justify-end gap-3">
        <Button variant="ghost" onClick={onClose} className="text-white hover:bg-white/10">
          Close
        </Button>
        <Button onClick={onClose}>
          {isUpgrade ? "Go to Stripe Checkout" : "Go to Stripe Portal"}
        </Button>
      </div>
    </Dialog>
  );
}
