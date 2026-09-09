import React, { useState } from "react";
import { Button, Input, Card, CardContent } from "@lh/ui";
import { useNavigate } from "react-router-dom";
import { isValidName, isValidPhone, formatPhone } from "@lh/utils";

export default function OnboardingPage() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    firstName: "",
    lastName: "",
    phoneNumber: "",
  });

  // Track validation errors natively in React to replace browser tooltips
  const [errors, setErrors] = useState<Record<string, string>>({});

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
    } else if (!isValidName(formData.firstName)) {
      newErrors.firstName = "Name can only contain letters, spaces, hyphens, and apostrophes";
    }
    
    if (!formData.lastName.trim()) {
      newErrors.lastName = "Last name is required";
    } else if (!isValidName(formData.lastName)) {
      newErrors.lastName = "Name can only contain letters, spaces, hyphens, and apostrophes";
    }
    
    if (!formData.phoneNumber.trim()) {
      newErrors.phoneNumber = "Phone number is required";
    } else if (!isValidPhone(formData.phoneNumber)) {
      newErrors.phoneNumber = "Please enter a valid 10-digit phone number";
    }
    
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault(); // Prevent page reload
    if (!validate()) return; // Stop if validation fails
    
    console.log("Saving user profile...", formData);
    // Mock routing to the dashboard
    navigate("/dashboard");
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
                <label className="flex items-center justify-between text-sm font-semibold text-white/80">
                  <span>Phone Number</span>
                  <span className="text-xs font-medium text-white/40 uppercase tracking-wider">Required for SMS</span>
                </label>
                <Input
                  name="phoneNumber"
                  type="tel"
                  value={formData.phoneNumber}
                  onChange={handleChange}
                  placeholder="(555) 000-0000"
                  error={errors.phoneNumber}
                  className="h-14 bg-white/5 px-4 text-base text-white placeholder:text-white/30 focus:ring-1 transition-colors border-white/10 focus:border-primary focus:ring-primary"
                />
              </div>

              <div className="pt-4">
                <Button 
                  type="submit" 
                  className="group relative h-14 w-full overflow-hidden rounded-lg bg-gradient-to-r from-primary to-red-800 text-base font-bold tracking-wide text-white shadow-lg shadow-primary/25 transition-all hover:scale-[1.02] hover:shadow-primary/40"
                >
                  <div className="absolute inset-0 flex h-full w-full justify-center [transform:skew(-12deg)_translateX(-100%)] group-hover:duration-1000 group-hover:[transform:skew(-12deg)_translateX(100%)]">
                    <div className="relative h-full w-8 bg-white/20" />
                  </div>
                  <span>Complete Setup</span>
                </Button>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
