import re

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'r') as f:
    content = f.read()

# First let's locate the 3 columns grid and replace it entirely to ensure no regex mess
start_marker = "{/* Social Auth Grid — 3 columns: Google, Facebook, Apple (coming soon) */}"
end_marker = "{/* Divider */}"

start_idx = content.find(start_marker)
end_idx = content.find(end_marker)

if start_idx != -1 and end_idx != -1:
    old_grid = content[start_idx:end_idx]
    
    new_grid = """{/* Social Auth Grid — 3 columns: Google, Facebook, Apple */}
            <div className="grid grid-cols-3 gap-3">
              {[ 
                { id: 'google', label: 'Google', icon: (
                  <svg className="h-4 w-4" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                    <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
                    <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
                    <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
                    <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
                  </svg>
                ), disabled: false },
                { id: 'facebook', label: 'Facebook', icon: (
                  <svg className="h-4 w-4" viewBox="0 0 24 24" fill="#1877F2" xmlns="http://www.w3.org/2000/svg">
                    <path d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.469h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.469h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"/>
                  </svg>
                ), disabled: false },
                { id: 'apple', label: 'Apple', icon: (
                  <svg className="h-4 w-4" viewBox="0 0 24 24" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                    <path d="M12.152 6.896c-.948 0-2.415-1.078-3.96-1.04-2.04.027-3.91 1.183-4.961 3.014-2.117 3.675-.546 9.103 1.519 12.09 1.013 1.454 2.208 3.09 3.792 3.039 1.52-.065 2.09-.987 3.935-.987 1.831 0 2.35.987 3.96.948 1.637-.026 2.62-1.458 3.579-2.91 1.163-1.696 1.64-3.342 1.66-3.428-.039-.013-3.182-1.222-3.22-4.857-.026-3.04 2.48-4.494 2.597-4.559-1.429-2.09-3.623-2.324-4.39-2.376-2-.156-3.675 1.09-4.502 1.09zM15.53 3.83c.843-1.012 1.4-2.427 1.245-3.83-1.207.052-2.662.805-3.532 1.818-.78.896-1.454 2.338-1.273 3.714 1.338.104 2.715-.688 3.56-1.702z"/>
                  </svg>
                ), disabled: true }
              ].map(provider => (
                <Button
                  key={provider.id}
                  type="button"
                  variant="outline"
                  aria-disabled={provider.disabled}
                  onClick={() => {
                    if (!provider.disabled) {
                      supabase.auth.signInWithOAuth({ provider: provider.id as any, options: { redirectTo: window.location.origin + "/dashboard" }});
                    }
                  }}
                  className={`group relative h-12 w-full border border-white/10 bg-black transition-all font-semibold shadow-sm overflow-hidden ${provider.disabled ? 'text-white/40 cursor-not-allowed' : 'text-white hover:bg-white/10'}`}
                >
                  <div className={`flex items-center justify-center gap-2 transition-transform duration-300 ${provider.disabled ? 'group-hover:-translate-y-12' : ''}`}>
                    {provider.icon}
                    {provider.label}
                  </div>
                  {provider.disabled && (
                    <div className="absolute inset-0 flex items-center justify-center translate-y-12 transition-transform duration-300 group-hover:translate-y-0 text-white/70">
                      Unavailable
                    </div>
                  )}
                </Button>
              ))}
            </div>

            """
    
    content = content.replace(old_grid, new_grid)

with open('apps/app/src/modules/auth/pages/LoginPage.tsx', 'w') as f:
    f.write(content)

