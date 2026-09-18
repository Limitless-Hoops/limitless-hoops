import * as React from "react"
import { Input, InputProps } from "./input"

export interface FormFieldProps extends InputProps {
  label: string;
  hint?: string;
}

const FormField = React.forwardRef<HTMLInputElement, FormFieldProps>(
  ({ label, hint, id, ...props }, ref) => {
    const fieldId = id ?? label.toLowerCase().replace(/\s+/g, "-");
    return (
      <div className="space-y-2">
        <label
          htmlFor={fieldId}
          className="text-xs font-semibold uppercase tracking-wider text-white/80"
        >
          {label}
        </label>
        <Input id={fieldId} ref={ref} {...props} />
        {hint && (
          <p className="text-xs text-white/40">{hint}</p>
        )}
      </div>
    );
  }
);
FormField.displayName = "FormField";

export { FormField };
