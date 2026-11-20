import re

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'r') as f:
    content = f.read()

# Fix isPhoneInput
old_is_phone = 'const isPhoneInput = /^\\d/.test(identifier);'
new_is_phone = 'const isPhoneInput = /^\\d|^\\(/.test(identifier);'
content = content.replace(old_is_phone, new_is_phone)

# Fix onChange
old_onchange = """                  onChange={(e) => {
                    const val = e.target.value;
                    setIdentifier(/^\\d/.test(val) ? formatPhone(val) : val);
                    if (error) setError("");
                  }}"""
new_onchange = """                  onChange={(e) => {
                    const val = e.target.value;
                    setIdentifier(/^\\d|^\\(/.test(val) ? formatPhone(val) : val);
                    if (error) setError("");
                  }}"""
content = content.replace(old_onchange, new_onchange)

# Wrap Google and Facebook contents in div (just in case it fixes the logo rendering issue in flex)
old_google = """              <Button
                variant="outline"
                className="h-12 w-full border border-white/10 bg-black text-white transition-all hover:bg-white/10 font-semibold shadow-sm flex items-center justify-center gap-2"
                onClick={() => supabase.auth.signInWithOAuth({ provider: 'google', options: { redirectTo: "http://localhost:5173/dashboard" }})}
              >
                <svg className="h-4 w-4" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
                  <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
                  <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
                  <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
                </svg>
                Google
              </Button>"""

new_google = """              <Button
                variant="outline"
                className="h-12 w-full border border-white/10 bg-black text-white transition-all hover:bg-white/10 font-semibold shadow-sm"
                onClick={() => supabase.auth.signInWithOAuth({ provider: 'google', options: { redirectTo: "http://localhost:5173/dashboard" }})}
              >
                <div className="flex items-center justify-center gap-2">
                  <svg className="h-4 w-4" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
                    <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
                    <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
                    <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
                  </svg>
                  Google
                </div>
              </Button>"""

content = content.replace(old_google, new_google)

old_facebook = """              <Button
                variant="outline"
                className="h-12 w-full border border-white/10 bg-black text-white transition-all hover:bg-white/10 font-semibold shadow-sm flex items-center justify-center gap-2"
                onClick={() => supabase.auth.signInWithOAuth({ provider: 'facebook', options: { redirectTo: "http://localhost:5173/dashboard" }})}
              >
                <svg className="h-4 w-4" viewBox="0 0 24 24" fill="#1877F2" xmlns="http://www.w3.org/2000/svg">
                  <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.469h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.469h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
                </svg>
                Facebook
              </Button>"""

new_facebook = """              <Button
                variant="outline"
                className="h-12 w-full border border-white/10 bg-black text-white transition-all hover:bg-white/10 font-semibold shadow-sm"
                onClick={() => supabase.auth.signInWithOAuth({ provider: 'facebook', options: { redirectTo: "http://localhost:5173/dashboard" }})}
              >
                <div className="flex items-center justify-center gap-2">
                  <svg className="h-4 w-4" viewBox="0 0 24 24" fill="#1877F2" xmlns="http://www.w3.org/2000/svg">
                    <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.469h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.469h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
                  </svg>
                  Facebook
                </div>
              </Button>"""

content = content.replace(old_facebook, new_facebook)

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'w') as f:
    f.write(content)

