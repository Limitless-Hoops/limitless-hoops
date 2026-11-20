import re

with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'r') as f:
    content = f.read()

# Fix phone formatting in HouseholdTab
old_phone_display = 'const phoneStr = user.phones?.find((p: any) => p.is_primary)?.phone || user.phone;'
new_phone_display = 'const phoneStr = formatPhone(user.phones?.find((p: any) => p.is_primary)?.phone || user.phone);'
content = content.replace(old_phone_display, new_phone_display)

# Fix URL persistence for activeTab
# First, add useSearchParams to imports
content = content.replace('import { useSearchParams, useNavigate } from "react-router-dom";', 'import { useSearchParams, useNavigate } from "react-router-dom";') # Wait, is it already imported?
if 'useSearchParams' not in content:
    content = content.replace('import { useNavigate, Link } from "react-router-dom";', 'import { useNavigate, Link, useSearchParams } from "react-router-dom";')
    # If neither exists:
    if 'useSearchParams' not in content:
        content = content.replace('import { supabase } from "@/lib/supabase";', 'import { useSearchParams } from "react-router-dom";\nimport { supabase } from "@/lib/supabase";')

old_active_tab = '  const [activeTab, setActiveTab] = useState("personal");'
new_active_tab = '  const [searchParams, setSearchParams] = useSearchParams();\n  const activeTab = searchParams.get("tab") || "personal";\n  const setActiveTab = (tab: string) => setSearchParams({ tab }, { replace: true });'
content = content.replace(old_active_tab, new_active_tab)


with open('apps/app/src/modules/account/pages/AccountPage.tsx', 'w') as f:
    f.write(content)

