"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import { ArrowRight, Lock, User, AlertCircle } from "lucide-react";

export default function LoginPage() {
  const router = useRouter();
  const [identifier, setIdentifier] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);

    try {
      const res = await fetch("http://localhost:8000/api/v1/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ identifier, password }),
      });

      const data = await res.json();

      if (!res.ok) {
        throw new Error(data.message || "Failed to login");
      }

      // Store token (in a real app, use secure cookie or httpOnly)
      localStorage.setItem("token", data.data.token);
      localStorage.setItem("user", JSON.stringify(data.data.user));
      
      router.push("/dashboard");
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-background flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <div className="flex justify-center items-center gap-2 mb-6">
          <div className="w-12 h-12 bg-primary rounded-xl flex items-center justify-center text-primary-foreground font-bold text-2xl">
            N
          </div>
        </div>
        <h2 className="mt-2 text-center text-3xl font-extrabold text-foreground tracking-tight">
          Masuk ke Akun Anda
        </h2>
        <p className="mt-2 text-center text-sm text-muted-foreground">
          Gunakan username atau email untuk masuk
        </p>
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
        <div className="glass py-8 px-4 shadow-2xl sm:rounded-2xl sm:px-10 border border-white/10">
          <form className="space-y-6" onSubmit={handleLogin}>
            {error && (
              <div className="bg-destructive/10 border border-destructive/20 text-destructive p-3 rounded-lg flex items-center gap-2 text-sm">
                <AlertCircle size={16} />
                {error}
              </div>
            )}
            
            <div>
              <label
                htmlFor="identifier"
                className="block text-sm font-medium text-foreground"
              >
                Username atau Email
              </label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <User className="h-5 w-5 text-muted-foreground" />
                </div>
                <input
                  id="identifier"
                  name="identifier"
                  type="text"
                  required
                  value={identifier}
                  onChange={(e) => setIdentifier(e.target.value)}
                  className="bg-muted/50 block w-full pl-10 sm:text-sm border-none rounded-xl py-3 focus:ring-1 ring-primary outline-none"
                  placeholder="admin@notopos.com atau admin"
                />
              </div>
            </div>

            <div>
              <label
                htmlFor="password"
                className="block text-sm font-medium text-foreground"
              >
                Password
              </label>
              <div className="mt-1 relative rounded-md shadow-sm">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Lock className="h-5 w-5 text-muted-foreground" />
                </div>
                <input
                  id="password"
                  name="password"
                  type="password"
                  required
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="bg-muted/50 block w-full pl-10 sm:text-sm border-none rounded-xl py-3 focus:ring-1 ring-primary outline-none"
                  placeholder="••••••••"
                />
              </div>
            </div>

            <div className="flex items-center justify-between">
              <div className="flex items-center">
                <input
                  id="remember-me"
                  name="remember-me"
                  type="checkbox"
                  className="h-4 w-4 text-primary focus:ring-primary border-gray-300 rounded"
                />
                <label
                  htmlFor="remember-me"
                  className="ml-2 block text-sm text-foreground"
                >
                  Ingat saya
                </label>
              </div>

              <div className="text-sm">
                <a
                  href="#"
                  className="font-medium text-primary hover:text-primary/80"
                >
                  Lupa password?
                </a>
              </div>
            </div>

            <div>
              <button
                type="submit"
                disabled={loading}
                className="w-full flex justify-center py-3 px-4 border border-transparent rounded-xl shadow-sm text-sm font-bold text-primary-foreground bg-primary hover:opacity-90 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary disabled:opacity-50 transition-all items-center gap-2"
              >
                {loading ? "Memproses..." : "Masuk"}
                {!loading && <ArrowRight size={16} />}
              </button>
            </div>
          </form>
        </div>
        
        {/* Helper info for testing */}
        <div className="mt-6 text-center text-xs text-muted-foreground bg-muted/20 p-4 rounded-xl border">
          <p className="font-bold mb-1">Testing Credentials (Mockup DB Offline):</p>
          <p>Email: <b>admin@notopos.com</b> OR Username: <b>admin</b></p>
          <p>Password: <b>password123</b></p>
        </div>
      </div>
    </div>
  );
}
