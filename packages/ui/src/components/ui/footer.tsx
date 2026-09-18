import LogoImage from "../../assets/lh-branding/logo-horizontal.svg"
import MottoImage from "../../assets/lh-branding/graphic-hero-motto.svg"

export function Footer() {
  return (
    <footer className="w-full border-t border-white/10 bg-black/50 py-8 px-4 sm:px-8 mt-auto">
      <div className="container mx-auto flex flex-col md:flex-row items-center justify-between gap-6">
        
        {/* Left side: Logo & Motto */}
        <div className="flex flex-col items-center md:items-start gap-4">
          <img 
            src={LogoImage} 
            alt="Limitless Hoops Logo" 
            className="h-10 object-contain brightness-0 invert opacity-80"
          />
          <img 
            src={MottoImage} 
            alt="Limitless Hoops Motto" 
            className="w-16 h-auto object-contain brightness-0 invert opacity-60"
          />
        </div>

        {/* Right side: Links & Copyright */}
        <div className="flex flex-col items-center md:items-end gap-3 text-sm text-white/50">
          <div className="flex items-center gap-4">
            <a href="/terms" className="hover:text-white transition-colors">Terms of Service</a>
            <span>&bull;</span>
            <a href="/privacy" className="hover:text-white transition-colors">Privacy Policy</a>
          </div>
          <p>&copy; {new Date().getFullYear()} Limitless Hoops. All rights reserved.</p>
        </div>
        
      </div>
    </footer>
  )
}
