import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import { AuthProvider } from "./modules/auth/contexts/AuthContext";
import { ProtectedRoute, PublicRoute, OnboardingRoute } from "./modules/auth/guards/AuthGuards";

import LoginPage from "./modules/auth/pages/LoginPage";
import VerifyPage from "./modules/auth/pages/VerifyPage";
import OnboardingPage from "./modules/auth/pages/OnboardingPage";
import NotFound from "./modules/static/NotFound";
import TermsOfServicePage from "./modules/static/TermsOfServicePage";
import PrivacyPolicyPage from "./modules/static/PrivacyPolicyPage";
import DashboardPage from "./modules/dashboard/pages/DashboardPage";
import AccountPage from "./modules/account/pages/AccountPage";

function App() {
  return (
    <AuthProvider>
      <Router>
        <Routes>
          {/* Default route redirect to login for now */}
          <Route path="/" element={<Navigate to="/login" replace />} />
          <Route path="*" element={<NotFound />} />
          
          {/* Legal */}
          <Route path="/terms" element={<TermsOfServicePage />} />
          <Route path="/privacy" element={<PrivacyPolicyPage />} />

          {/* Auth Module */}
          <Route path="/login" element={
            <PublicRoute>
              <LoginPage />
            </PublicRoute>
          } />
          <Route path="/auth/verify" element={
            <PublicRoute>
              <VerifyPage />
            </PublicRoute>
          } />
          <Route path="/auth/onboarding" element={
            <OnboardingRoute>
              <OnboardingPage />
            </OnboardingRoute>
          } />
          
          {/* Dashboard (Protected) */}
          <Route path="/dashboard" element={
            <ProtectedRoute>
              <DashboardPage />
            </ProtectedRoute>
          } />
          <Route path="/account" element={
            <ProtectedRoute>
              <AccountPage />
            </ProtectedRoute>
          } />
        </Routes>
      </Router>
    </AuthProvider>
  );
}

export default App;
