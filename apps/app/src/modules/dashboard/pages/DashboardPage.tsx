import DashboardLayout from "../components/DashboardLayout";
import { Button } from "@lh/ui";
import { useNavigate } from "react-router-dom";

export default function DashboardPage() {
  const navigate = useNavigate();

  return (
    <DashboardLayout>
      <div className="flex flex-col gap-8">
        <div className="flex flex-col gap-2">
          <h1 className="text-4xl font-extrabold tracking-tight text-white uppercase">Welcome, Jon</h1>
          <p className="text-white/60 text-lg">Manage your household and subscription.</p>
        </div>
        
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {/* Account Management Card */}
          <div className="flex min-h-48 flex-col justify-between rounded-2xl border border-white/10 bg-white/5 p-6 shadow-xl backdrop-blur-sm">
            <div>
              <h2 className="text-lg font-bold uppercase tracking-widest text-white mb-2">
                Household & Account
              </h2>
              <p className="text-sm text-white/50 leading-relaxed mb-6">
                View your current subscription plan, manage household members, and update your personal information.
              </p>
            </div>
            
            <Button 
              variant="outline" 
              className="w-full border-white/20 text-white hover:bg-white/10 font-semibold tracking-wide"
              onClick={() => navigate("/account")}
            >
              Manage Account
            </Button>
          </div>
        </div>
      </div>
    </DashboardLayout>
  );
}
