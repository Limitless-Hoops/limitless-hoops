import * as React from "react"

export interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  error?: string;
  prefixText?: string;
}

const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ className = "", type, error, prefixText, ...props }, ref) => {
    // Removed 'flex' to fix native input elements like type="date"
    const defaultClasses = `w-full rounded-xl border bg-white/5 ${prefixText ? "pl-10 pr-4" : "px-4"} h-12 text-sm text-white shadow-sm transition-all hover:bg-white/10 file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-white/30 focus-visible:outline-none focus-visible:ring-2 disabled:cursor-not-allowed disabled:opacity-50`;
    
    // Allow error to override border and ring colors
    const stateClasses = error 
      ? "border-destructive focus-visible:border-destructive focus-visible:ring-destructive/50" 
      : "border-white/10 focus-visible:border-primary focus-visible:ring-primary/50 focus-visible:bg-white/10";

    return (
      <div className="w-full flex flex-col gap-2 relative">
        {prefixText && (
          <div className="absolute left-0 top-0 bottom-0 flex items-center pl-4 pointer-events-none text-white/50 text-sm">
            {prefixText}
          </div>
        )}
        <input
          type={type}
          className={`${defaultClasses} ${stateClasses} ${className}`}
          ref={ref}
          {...props}
        />
        {error && (
          <p className="text-xs font-medium text-destructive animate-in slide-in-from-top-1">
            {error}
          </p>
        )}
      </div>
    )
  }
)
Input.displayName = "Input"

export { Input }
