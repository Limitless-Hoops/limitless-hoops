import re

with open('packages/ui/src/components/ui/form-field.tsx', 'r') as f:
    content = f.read()

content = content.replace('<Input id={fieldId} ref={ref} {...props} />', '<Input id={fieldId} ref={ref} {...props} prefixText={props.prefixText} />')
# Actually wait, {...props} already spreads prefixText! I don't need to do anything to FormField.tsx!

