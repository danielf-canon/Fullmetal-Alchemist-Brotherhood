// src/App.tsx
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { AuthProvider } from "./context/AuthContext";
import ProtectedRoute from "./components/layout/ProtectedRoute";

import Navbar from "./components/layout/Navbar";
import Sidebar from "./components/layout/Sidebar";

import Login from "./pages/Login";
import Register from "./pages/Register";

import Dashboard from "./pages/Dashboard";
import Alquimistas from "./pages/Alquimistas";
import Materials from "./pages/Materials";
import Missions from "./pages/Missions";
import Transmutations from "./pages/Transmutations";
import Auditorias from "./pages/Auditorias";

export default function App() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <div style={{ display: "flex", minHeight: "100vh" }}>
          <Sidebar />
            
          <div style={{ flex: 1 }}>
            <Navbar />

            <div style={{ padding: 20 }}>
              <Routes>
                {/* RUTAS PÚBLICAS */}
                <Route path="/login" element={<Login />} />
               
                <Route path="/register" element={<Register />} />

                {/* RUTAS PROTEGIDAS */}
                <Route
                  path="/"
                  element={
                    <ProtectedRoute>
                      <Dashboard />
                    </ProtectedRoute>
                  }
                />

                <Route
                  path="/alquimistas"
                  element={
                    <ProtectedRoute>
                      <Alquimistas />
                    </ProtectedRoute>
                  }
                />

                <Route
                  path="/materials"
                  element={
                    <ProtectedRoute>
                      <Materials />
                    </ProtectedRoute>
                  }
                />

                <Route
                  path="/missions"
                  element={
                    <ProtectedRoute>
                      <Missions />
                    </ProtectedRoute>
                  }
                />

                <Route
                  path="/transmutations"
                  element={
                    <ProtectedRoute>
                      <Transmutations />
                    </ProtectedRoute>
                  }
                />

                <Route
                  path="/auditorias"
                  element={
                    <ProtectedRoute>
                      <Auditorias />
                    </ProtectedRoute>
                  }
                />
              </Routes>
            </div>
          </div>
        </div>
      </AuthProvider>
    </BrowserRouter>
  );
}
