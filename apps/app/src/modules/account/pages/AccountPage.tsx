import { useState } from "react";
import DashboardLayout from "../../dashboard/components/DashboardLayout";
import { Card, CardContent, Button, FormField } from "@lh/ui";
import { formatPhone, isValidName } from "@lh/utils";

// ─── Constants ───────────────────────────────────────────────────────────────

const TABS = [
  { id: "personal", label: "Personal Information" },
  { id: "security", label: "Login & Security" },
  { id: "household", label: "Household & Contacts" },
  { id: "billing", label: "Subscription & Billing" },
  { id: "advanced", label: "Advanced" },
];

/** Shared card class applied to every section card */
const cardClass = "border-white/10 bg-black/40 backdrop-blur-2xl shadow-xl shadow-black/50";

// ─── Tabs ─────────────────────────────────────────────────────────────────────

function PersonalInfoTab() {
  const [firstName, setFirstName] = useState("Jon");
  const [lastName, setLastName] = useState("Sanders");
  const [phone, setPhone] = useState("(555) 123-4567");

  const firstNameError = firstName && !isValidName(firstName) ? "Name can only contain letters, spaces, hyphens, and apostrophes." : undefined;
  const lastNameError = lastName && !isValidName(lastName) ? "Name can only contain letters, spaces, hyphens, and apostrophes." : undefined;

  const handlePhoneChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPhone(formatPhone(e.target.value));
  };

  return (
    <Card className={cardClass}>
      <CardContent className="p-6 sm:p-8 space-y-6">
        <div>
          <h2 className="text-xl font-bold text-white mb-1">Personal Information</h2>
          <p className="text-sm text-white/50">Manage your basic profile details.</p>
        </div>

        <div className="grid gap-6 sm:grid-cols-2">
          <FormField
            label="First Name"
            value={firstName}
            error={firstNameError}
            onChange={(e) => setFirstName(e.target.value)}
            className="bg-white/5 border-white/10 text-white"
          />
          <FormField
            label="Last Name"
            value={lastName}
            error={lastNameError}
            onChange={(e) => setLastName(e.target.value)}
            className="bg-white/5 border-white/10 text-white"
          />
        </div>

        <FormField
          label="Date of Birth"
          type="date"
          defaultValue="1980-01-01"
          hint="Used to verify you are 18 or older to sign legal waivers and consent forms on behalf of your dependents."
          className="bg-white/5 border-white/10 text-white w-full sm:w-1/2"
        />

        <FormField
          label="Primary Phone Number"
          type="tel"
          value={phone}
          onChange={handlePhoneChange}
          className="bg-white/5 border-white/10 text-white w-full sm:w-1/2"
        />

        <div className="pt-4 flex justify-end">
          <Button className="bg-primary text-white hover:bg-primary/80">Save Changes</Button>
        </div>
      </CardContent>
    </Card>
  );
}

function SecurityTab() {
  return (
    <Card className={cardClass}>
      <CardContent className="p-6 sm:p-8 space-y-8">
        <div>
          <h2 className="text-xl font-bold text-white mb-1">Login & Security</h2>
          <p className="text-sm text-white/50">Update your email and manage connected accounts.</p>
        </div>

        <div className="space-y-4">
          <h3 className="text-sm font-semibold uppercase tracking-wider text-white/80 border-b border-white/10 pb-2">Email Address</h3>
          <div className="flex flex-col sm:flex-row gap-4 items-start sm:items-end">
            <div className="w-full sm:w-2/3">
              <FormField
                label="Email"
                type="email"
                defaultValue="jon.sanders@example.com"
                className="bg-white/5 border-white/10 text-white"
              />
            </div>
            <Button variant="outline" className="border-white/20 hover:bg-white/10 whitespace-nowrap">Verify New Email</Button>
          </div>
          <p className="text-xs text-white/40">You must verify your new email address before it takes effect.</p>
        </div>

        <div className="space-y-4">
          <h3 className="text-sm font-semibold uppercase tracking-wider text-white/80 border-b border-white/10 pb-2">Linked Accounts</h3>
          <div className="flex items-center justify-between bg-white/5 border border-white/10 rounded-lg p-4">
            <div className="flex items-center gap-4">
              <div className="bg-white p-1.5 rounded-full">
                <svg className="h-5 w-5" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
                  <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
                  <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
                  <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
                </svg>
              </div>
              <div>
                <p className="font-semibold text-white">Google</p>
                <p className="text-xs text-white/50">Connected as jon.sanders@example.com</p>
              </div>
            </div>
            <Button variant="outline" className="border-white/20 hover:bg-red-900/30 hover:text-red-400 hover:border-red-900/50" disabled>Disconnect</Button>
          </div>
          <p className="text-xs text-white/40">You cannot disconnect Google because it is your only login method.</p>
        </div>
      </CardContent>
    </Card>
  );
}

function HouseholdTab() {
  const [guardianPhone, setGuardianPhone] = useState("(555) 987-6543");

  return (
    <div className="space-y-8">
      {/* Additional Contacts */}
      <Card className={cardClass}>
        <CardContent className="p-6 sm:p-8 space-y-6">
          <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
            <div>
              <h2 className="text-xl font-bold text-white mb-1">Additional Guardians</h2>
              <p className="text-sm text-white/50">Add co-parents or guardians to receive notifications.</p>
            </div>
            <Button variant="outline" className="border-white/20 hover:bg-white/10">+ Add Guardian</Button>
          </div>

          <div className="bg-white/5 border border-white/10 rounded-lg p-4 flex justify-between items-center">
            <div>
              <p className="font-semibold text-white">Jane Smith</p>
              <p className="text-xs text-white/50">jane.smith@example.com &bull; {guardianPhone}</p>
            </div>
            <Button variant="ghost" className="text-white/40 hover:text-red-400 hover:bg-red-900/20">Remove</Button>
          </div>
        </CardContent>
      </Card>

      {/* Dependents */}
      <Card className={cardClass}>
        <CardContent className="p-6 sm:p-8 space-y-6">
          <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
            <div>
              <h2 className="text-xl font-bold text-white mb-1">Dependents</h2>
              <p className="text-sm text-white/50">Manage the players in your household.</p>
            </div>
            <Button className="bg-primary text-white hover:bg-primary/80">+ Add Player</Button>
          </div>

          {[
            { name: "Jimmy Smith", age: 12, dob: "Mar 14, 2014" },
            { name: "Sarah Smith", age: 15, dob: "Jun 22, 2011" },
          ].map((dep) => (
            <div key={dep.name} className="bg-white/5 border border-white/10 rounded-lg p-4 flex justify-between items-center">
              <div>
                <p className="font-semibold text-white">{dep.name}</p>
                <p className="text-xs text-white/50">Age {dep.age} &bull; Born {dep.dob}</p>
              </div>
              <Button variant="ghost" className="text-white/40 hover:text-red-400 hover:bg-red-900/20">Remove</Button>
            </div>
          ))}
        </CardContent>
      </Card>
    </div>
  );
}

function BillingTab() {
  return (
    <Card className={cardClass}>
      <CardContent className="p-6 sm:p-8 space-y-6">
        <div>
          <h2 className="text-xl font-bold text-white mb-1">Subscription & Billing</h2>
          <p className="text-sm text-white/50">Manage your membership plan and payment methods.</p>
        </div>

        <div className="bg-linear-to-r from-primary/20 to-transparent border border-primary/30 rounded-lg p-6">
          <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-6">
            <div>
              <h3 className="text-sm font-bold uppercase tracking-widest text-primary mb-1">Current Plan</h3>
              <p className="text-2xl font-extrabold text-white">Limitless Member</p>
              <p className="text-sm text-white/60 mt-1">$49.00 / month &bull; Renews on Oct 1, 2026</p>
            </div>
            <Button variant="outline" className="w-full md:w-auto border-white/20 hover:bg-white/10">Upgrade Plan</Button>
          </div>
        </div>

        <div className="pt-4 border-t border-white/10 flex justify-end">
          <Button variant="ghost" className="text-white/40 hover:text-white hover:bg-white/5">Cancel Subscription</Button>
        </div>
      </CardContent>
    </Card>
  );
}

function AdvancedTab() {
  return (
    <Card className="border-red-900/30 bg-black/40 backdrop-blur-2xl shadow-xl shadow-black/50">
      <CardContent className="p-6 sm:p-8 space-y-6">
        <div>
          <h2 className="text-xl font-bold text-red-500 mb-1">Danger Zone</h2>
          <p className="text-sm text-white/50">Irreversible and destructive actions.</p>
        </div>

        <div className="border border-red-900/50 bg-red-950/20 rounded-lg p-6 flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
          <div>
            <h3 className="font-semibold text-white">Delete Account</h3>
            <p className="text-xs text-white/50 mt-1 max-w-md">
              This action will permanently delete your account and archive all dependent data. You will immediately lose access to the platform.
            </p>
          </div>
          <Button className="bg-red-600 text-white hover:bg-red-700 whitespace-nowrap shrink-0">Delete Account</Button>
        </div>
      </CardContent>
    </Card>
  );
}

// ─── Page ─────────────────────────────────────────────────────────────────────

export default function AccountPage() {
  const [activeTab, setActiveTab] = useState("personal");

  return (
    <DashboardLayout>
      <div className="flex flex-col lg:flex-row gap-8 pb-12">

        {/* Sidebar Nav */}
        <aside className="w-full lg:w-64 shrink-0">
          <div className="sticky top-24">
            <h1 className="text-2xl font-bold tracking-tight text-white uppercase mb-6 hidden lg:block">Account Settings</h1>

            <nav className="flex lg:flex-col gap-2 overflow-x-auto pb-4 lg:pb-0">
              {TABS.map((tab) => (
                <button
                  key={tab.id}
                  onClick={() => setActiveTab(tab.id)}
                  className={`whitespace-nowrap rounded-lg px-4 py-3 text-left text-sm font-medium transition-colors ${
                    activeTab === tab.id
                      ? "bg-primary/20 text-white border border-primary/50"
                      : "text-white/60 hover:bg-white/5 hover:text-white border border-transparent"
                  }`}
                >
                  {tab.label}
                </button>
              ))}
            </nav>
          </div>
        </aside>

        {/* Main Content Area */}
        <div className="flex-1 min-w-0">
          <h1 className="text-2xl font-bold tracking-tight text-white uppercase mb-6 lg:hidden">Account Settings</h1>

          <div key={activeTab} className="animate-in fade-in slide-in-from-bottom-4 duration-300">
            {activeTab === "personal"  && <PersonalInfoTab />}
            {activeTab === "security"  && <SecurityTab />}
            {activeTab === "household" && <HouseholdTab />}
            {activeTab === "billing"   && <BillingTab />}
            {activeTab === "advanced"  && <AdvancedTab />}
          </div>
        </div>

      </div>
    </DashboardLayout>
  );
}
