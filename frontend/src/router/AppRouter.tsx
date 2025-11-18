import { BrowserRouter, Routes, Route } from "react-router-dom";

import ProtectedRoute from "../components/layout/ProtectedRoute";

import Login from "../pages/Login";
import Register from "../pages/Register";
import Dashboard from "../pages/Dashboard";

import AlquimistasPage from "../pages/Alquimistas";
import MaterialsPage from "../pages/Materials";
import MissionsPage from "../pages/Missions";
import TransmutationsPage from "../pages/Transmutations";
import AuditoriasPage from "../pages/Auditorias";

export default function AppRouter() {
  return (
    <BrowserRouter>
      <Routes>

        {/* Public routes */}
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />

        {/* Protected */}
        <Route path="/" element={
          <ProtectedRoute>
            <Dashboard />
          </ProtectedRoute>
        } />

        <Route path="/alquimistas" element={
          <ProtectedRoute>
            <AlquimistasPage />
          </ProtectedRoute>
        } />

        <Route path="/materials" element={
          <ProtectedRoute>
            <MaterialsPage />
          </ProtectedRoute>
        } />

        <Route path="/missions" element={
          <ProtectedRoute>
            <MissionsPage />
          </ProtectedRoute>
        } />

        <Route path="/transmutations" element={
          <ProtectedRoute>
            <TransmutationsPage />
          </ProtectedRoute>
        } />

        <Route path="/auditorias" element={
          <ProtectedRoute>
            <AuditoriasPage />
          </ProtectedRoute>
        } />
      </Routes>
    </BrowserRouter>
  );
}
