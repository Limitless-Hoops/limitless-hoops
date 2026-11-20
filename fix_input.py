import re

with open('packages/ui/src/components/ui/input.tsx', 'r') as f:
    content = f.read()

old_interface = """export interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  error?: string;
}"""
new_interface = """export interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  error?: string;
  prefixText?: string;
}"""
content = content.replace(old_interface, new_interface)

old_comp = """const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ className = "", type, error, ...props }, ref) => {
    // Removed 'flex' to fix native input elements like type="date"
    const defaultClasses = "w-full rounded-xl border bg-white/5 px-4 h-12 text-sm text-white shadow-sm transition-all hover:bg-white/10 file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-white/30 focus-visible:outline-none focus-visible:ring-2 disabled:cursor-not-allowed disabled:opacity-50";"""

new_comp = """const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({ className = "", type, error, prefixText, ...props }, ref) => {
    // Removed 'flex' to fix native input elements like type="date"
    const defaultClasses = `w-full rounded-xl border bg-white/5 ${prefixText ? "pl-10 pr-4" : "px-4"} h-12 text-sm text-white shadow-sm transition-all hover:bg-white/10 file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-white/30 focus-visible:outline-none focus-visible:ring-2 disabled:cursor-not-allowed disabled:opacity-50`;"""

content = content.replace(old_comp, new_comp)

old_return = """    return (
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
    )"""

new_return = """    return (
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
    )"""

content = content.replace(old_return, new_return)

with open('packages/ui/src/components/ui/input.tsx', 'w') as f:
    f.write(content)

