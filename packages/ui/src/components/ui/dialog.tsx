import * as React from "react";
import { useEffect, useRef } from "react";

export interface DialogProps extends React.DialogHTMLAttributes<HTMLDialogElement> {
  isOpen: boolean;
  onClose: () => void;
  title?: string;
  description?: string;
}

const Dialog = React.forwardRef<HTMLDialogElement, DialogProps>(
  ({ isOpen, onClose, title, description, children, className = "", ...props }, ref) => {
    const internalRef = useRef<HTMLDialogElement>(null);
    const dialogRef = (ref as any) || internalRef;

    useEffect(() => {
      const dialog = dialogRef.current;
      if (!dialog) return;

      if (isOpen && !dialog.open) {
        dialog.showModal();
      } else if (!isOpen && dialog.open) {
        dialog.close();
      }
    }, [isOpen, dialogRef]);

    useEffect(() => {
      const dialog = dialogRef.current;
      if (!dialog) return;

      const handleCancel = (e: Event) => {
        e.preventDefault();
        onClose();
      };

      dialog.addEventListener("cancel", handleCancel);
      return () => dialog.removeEventListener("cancel", handleCancel);
    }, [onClose, dialogRef]);

    const handleBackdropClick = (e: React.MouseEvent<HTMLDialogElement>) => {
      if (e.target === e.currentTarget) {
        onClose();
      }
    };

    return (
      <dialog
        ref={dialogRef}
        onClick={handleBackdropClick}
        className={`fixed m-auto rounded-2xl border border-white/10 bg-black/60 p-0 text-white shadow-2xl backdrop-blur-2xl backdrop:bg-black/60 backdrop:backdrop-blur-sm open:animate-in open:fade-in open:zoom-in-95 closed:animate-out closed:fade-out closed:zoom-out-95 ${className}`}
        {...props}
      >
        <div className="w-full max-w-lg p-6 sm:p-8">
          {(title || description) && (
            <div className="mb-6">
              {title && <h2 className="text-2xl font-bold tracking-tight text-white mb-2">{title}</h2>}
              {description && <p className="text-sm text-white/60">{description}</p>}
            </div>
          )}
          {children}
        </div>
      </dialog>
    );
  }
);
Dialog.displayName = "Dialog";

export { Dialog };
