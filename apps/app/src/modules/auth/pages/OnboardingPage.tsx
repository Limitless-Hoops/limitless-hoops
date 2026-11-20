import React, { useState, useEffect } from "react";
import { Button, Input, Card, CardContent } from "@lh/ui";
import { useNavigate } from "react-router-dom";
import { formatPhone, normalizePhoneForAPI, getDateOfBirthMax } from "@lh/utils";
import { supabase } from "@/lib/supabase";
import { useAuth } from "@/modules/auth/contexts/AuthContext";

export default function OnboardingPage() {
  const navigate = useNavigate();
  const { session, refreshHousehold, signOut } = useAuth();
  const [formData, setFormData] = useState({
    firstName: "",
    lastName: "",
    email: "", // Added for backend integration
    phoneNumber: "",
    dob: "",
  });

  // Pre-fill data from Google OAuth session
  useEffect(() => {
    if (session?.user) {
      const metadata = session.user.user_metadata || {};
      const fullName = metadata.full_name || metadata.name || "";
      const userEmail = session.user.email || "";
      const userPhone = session.user.phone || "";
      
      let initialFirstName = "";
      let initialLastName = "";
      
      if (fullName) {
        const parts = fullName.trim().split(" ");
        initialFirstName = parts[0];
        initialLastName = parts.slice(1).join(" ");
      }

      setFormData(prev => ({
        ...prev,
        firstName: prev.firstName || initialFirstName,
        lastName: prev.lastName || initialLastName,
        email: prev.email || userEmail,
        phoneNumber: prev.phoneNumber || (userPhone ? formatPhone(userPhone) : ""),
      }));
    }
  }, [session]);

  // Track validation errors natively in React to replace browser tooltips
  const [errors, setErrors] = useState<Record<string, string>>({});
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [apiError, setApiError] = useState("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    
    let parsedValue = value;
    if (name === "phoneNumber") {
      parsedValue = formatPhone(value);
    }
    
    setFormData((prev) => ({ ...prev, [name]: parsedValue }));
    // Clear error for the field once the user starts typing
    if (errors[name]) {
      setErrors((prev) => ({ ...prev, [name]: "" }));
    }
  };

  const validate = () => {
    const newErrors: Record<string, string> = {};
    if (!formData.firstName.trim()) {
      newErrors.firstName = "First name is required";
    }
    if (!formData.lastName.trim()) {
      newErrors.lastName = "Last name is required";
    }
    // C6: Email is required UNLESS the user authenticated via phone OTP
    // (in which case session.user.email will be empty and phone is their identity)
    const isPhoneOnlyUser = !session?.user?.email && !!session?.user?.phone;
    if (!isPhoneOnlyUser) {
      if (!formData.email.trim() || !formData.email.includes("@")) {
        newErrors.email = "Valid email address is required";
      }
    }
    if (!formData.phoneNumber.trim()) {
      newErrors.phoneNumber = "Phone number is required";
    }
    if (!formData.dob.trim()) {
      newErrors.dob = "Date of Birth is required";
    }
    
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); // Prevent page reload
    if (!validate()) return; // Stop if validation fails
    
    setIsSubmitting(true);
    setApiError("");

    try {
      const { data: { session } } = await supabase.auth.getSession();
      if (!session) {
        throw new Error("No active session. Please log in again.");
      }

      const response = await fetch("http://localhost:3000/api/households", {
        method: "POST",
        headers: { 
          "Content-Type": "application/json",
          "Authorization": `Bearer ${session.access_token}`
        },
        body: JSON.stringify({
          parent: {
            first_name: formData.firstName,
            last_name: formData.lastName,
            email: formData.email,
            phone: normalizePhoneForAPI(formData.phoneNumber),
            dob: formData.dob,
          },
          players: [], // We create the household first, they add players later in the dashboard
        }),
      });

      if (!response.ok) {
        let errMsg = "Failed to create household on the server";
        try {
          const resJson = await response.json();
          if (resJson.error) errMsg = resJson.error;
        } catch (e) {
          // Ignore JSON parsing errors
        }
        throw new Error(errMsg);
      }

      await refreshHousehold();
      navigate("/dashboard");
    } catch (err: any) {
      console.error(err);
      setApiError(err.message || "Something went wrong. Please ensure the API is running.");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="relative flex min-h-screen flex-col items-center justify-center overflow-hidden bg-background p-4 sm:p-8">
      {/* Subtle Premium Background Glow */}
      <div className="absolute top-[-10%] left-1/2 h-[500px] w-[800px] -translate-x-1/2 rounded-[100%] bg-primary/20 opacity-50 blur-[120px]" />

      <div className="z-10 flex w-full max-w-md flex-col items-center">
        <div className="mb-10 text-center flex flex-col items-center animate-in fade-in slide-in-from-bottom-4 duration-700">
          <h1 className="text-3xl font-extrabold uppercase tracking-tight text-white sm:text-4xl">
            Complete Your Profile
          </h1>
          <p className="mt-3 text-sm text-white/60 sm:text-base font-medium max-w-sm">
            Just a few details before you manage your household and secure a roster spot.
          </p>
        </div>

        {/* Glassmorphic Card */}
        <Card className="w-full border-white/10 bg-black/40 backdrop-blur-2xl shadow-2xl shadow-black/50 animate-in fade-in slide-in-from-bottom-6 duration-1000 delay-150 fill-mode-both">
          <CardContent className="p-6 sm:p-8">
            {/* noValidate disables the ugly native browser popups */}
            <form onSubmit={handleSubmit} noValidate className="space-y-6">
              
              <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
                <div className="space-y-2">
                  <label className="text-sm font-semibold text-white/80">
                    First Name
                  </label>
                  <Input
                    name="firstName"
                    value={formData.firstName}
                    onChange={handleChange}
                    placeholder="e.g. John"
                    error={errors.firstName}
                    className="h-14 bg-white/5 px-4 text-base text-white placeholder:text-white/30 focus:ring-1 transition-colors border-white/10 focus:border-primary focus:ring-primary"
                  />
                </div>

                <div className="space-y-2">
                  <label className="text-sm font-semibold text-white/80">
                    Last Name
                  </label>
                  <Input
                    name="lastName"
                    value={formData.lastName}
                    onChange={handleChange}
                    placeholder="e.g. Doe"
                    error={errors.lastName}
                    className="h-14 bg-white/5 px-4 text-base text-white placeholder:text-white/30 focus:ring-1 transition-colors border-white/10 focus:border-primary focus:ring-primary"
                  />
                </div>
              </div>

              <div className="space-y-2">
                <label className="text-sm font-semibold text-white/80">
                  Email Address
                </label>
                <Input
                  name="email"
                  type="email"
                  value={formData.email}
                  onChange={handleChange}
                  placeholder="parent@example.com"
                  error={errors.email}
                  disabled={!!session?.user?.email}
                  className="h-14 bg-white/5 px-4 text-base text-white placeholder:text-white/30 focus:ring-1 transition-colors border-white/10 focus:border-primary focus:ring-primary disabled:opacity-50 disabled:cursor-not-allowed"
                />
              </div>

              <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
                <div className="space-y-2">
                  <label className="flex items-center justify-between text-sm font-semibold text-white/80">
                    <span>Phone Number</span>
                    <span className="text-xs font-medium text-white/40 uppercase tracking-wider">Required for SMS</span>
                  </label>
                  <Input
                    name="phoneNumber"
                    type="tel" prefixText="+1"
                    value={formData.phoneNumber}
                    onChange={handleChange}
                    placeholder="(555) 000-0000"
                    error={errors.phoneNumber}
                    disabled={!!session?.user?.phone}
                    className="h-14 bg-white/5 px-4 text-base text-white placeholder:text-white/30 focus:ring-1 transition-colors border-white/10 focus:border-primary focus:ring-primary disabled:opacity-50 disabled:cursor-not-allowed"
                  />
                </div>

                <div className="space-y-2">
                  <label className="text-sm font-semibold text-white/80">
                    Date of Birth
                  </label>
                  <Input
                    name="dob"
                    type="date"
                    value={formData.dob}
                    max={getDateOfBirthMax(18)}
                    onChange={handleChange}
                    error={errors.dob}
                    className="h-14 bg-white/5 px-4 text-base text-white placeholder:text-white/30 focus:ring-1 transition-colors border-white/10 focus:border-primary focus:ring-primary"
                  />
                </div>
              </div>

              {apiError && (
                <div className="text-red-500 text-sm text-center font-medium">
                  {apiError}
                </div>
              )}

              <div className="pt-4">
                <Button 
                  type="submit" 
                  disabled={isSubmitting}
                  className="group relative h-14 w-full overflow-hidden rounded-lg bg-gradient-to-r from-primary to-red-800 text-base font-bold tracking-wide text-white shadow-lg shadow-primary/25 transition-all hover:scale-[1.02] hover:shadow-primary/40 disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  <div className="absolute inset-0 flex h-full w-full justify-center [transform:skew(-12deg)_translateX(-100%)] group-hover:duration-1000 group-hover:[transform:skew(-12deg)_translateX(100%)]">
                    <div className="relative h-full w-8 bg-white/20" />
                  </div>
                  <span>Complete Setup</span>
                </Button>
              </div>

              <div className="text-center pt-2">
                <button
                  type="button"
                  onClick={() => signOut().then(() => navigate("/login", { replace: true }))}
                  className="cursor-pointer text-xs font-medium text-white/60 underline underline-offset-2 transition-colors hover:text-white"
                >
                  Sign out and use a different account
                </button>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
