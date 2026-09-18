import { Button } from "@lh/ui";
import { useNavigate } from "react-router-dom";
import LogoImage from "@lh/ui/assets/lh-branding/logo-horizontal.svg";

export default function NotFound() {
  const navigate = useNavigate();

  return (
    <div className="relative flex min-h-screen flex-col items-center justify-center overflow-hidden bg-background p-4 sm:p-8">
      {/* Subtle Premium Background Glow */}
      <div className="absolute top-[-10%] left-1/2 h-125 w-200 -translate-x-1/2 rounded-[100%] bg-primary/20 opacity-40 blur-[120px]" />

      <div className="z-10 flex w-full max-w-md flex-col items-center text-center">
        <img 
          src={LogoImage} 
          alt="Limitless Hoops Logo" 
          className="mb-12 w-40 object-contain drop-shadow-2xl sm:w-48 brightness-0 invert opacity-50 hover:opacity-100 transition-opacity duration-500 cursor-pointer"
          onClick={() => navigate("/")}
        />

        <div className="relative mb-6 animate-bounce">
          <div className="text-8xl sm:text-[120px] drop-shadow-[0_0_40px_rgba(220,38,38,0.4)]">
            🏀
          </div>
        </div>

        <h1 className="text-5xl font-extrabold uppercase tracking-tight text-white sm:text-7xl drop-shadow-md mb-2 animate-in fade-in slide-in-from-bottom-4 duration-700">
          Airball!
        </h1>
        
        <h2 className="text-xl font-bold mb-6 uppercase tracking-widest text-white/80 animate-in fade-in slide-in-from-bottom-6 duration-700 delay-100 fill-mode-both">
          404 - Out of Bounds
        </h2>

        <p className="mb-10 text-sm text-white/60 sm:text-base font-medium max-w-sm leading-relaxed animate-in fade-in slide-in-from-bottom-8 duration-700 delay-200 fill-mode-both">
          Looks like you stepped out of bounds or the page you're looking for got traded. Let's get you back in the game.
        </p>

        <Button
          onClick={() => navigate("/")}
          className="group relative h-14 w-full sm:w-3/4 overflow-hidden rounded-lg bg-linear-to-r from-primary to-red-800 text-base font-bold tracking-wide text-white shadow-lg shadow-primary/25 transition-all hover:scale-[1.05] hover:shadow-primary/40 animate-in fade-in zoom-in duration-700 delay-300 fill-mode-both"
        >
          <div className="absolute inset-0 flex h-full w-full justify-center transform-[skew(-12deg)_translateX(-100%)] group-hover:duration-1000 group-hover:transform-[skew(-12deg)_translateX(100%)]">
            <div className="relative h-full w-8 bg-white/20" />
          </div>
          <span>Return to Court</span>
        </Button>      
      </div>
    </div>
  );
}
