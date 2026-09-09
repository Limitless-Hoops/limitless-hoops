import * as React from "react"

export interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  error?: string;
}

const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ className = "", type, error, ...props }, ref) => {
    const defaultClasses = "flex h-9 w-full rounded-md border bg-transparent px-3 py-1 text-sm shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 disabled:cursor-not-allowed disabled:opacity-50";
    
    // Allow error to override border and ring colors
    const stateClasses = error 
      ? "border-destructive focus:border-destructive focus-visible:border-destructive focus-visible:ring-destructive/50" 
      : "border-input focus-visible:ring-ring";

    return (
      <div className="w-full flex flex-col gap-2">
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
