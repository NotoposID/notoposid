"use client";

import React from "react";
import {
  LayoutDashboard,
  ShoppingCart,
  Package,
  Users,
  Settings,
  PieChart,
  Bell,
  Search,
  Plus
} from "lucide-react";

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className="flex h-screen bg-muted/20">
      {/* Sidebar */}
      <aside className="w-64 bg-background border-r flex flex-col">
        <div className="p-6 flex items-center gap-2 border-b">
          <div className="w-8 h-8 bg-primary rounded-lg flex items-center justify-center text-primary-foreground font-bold">
            N
          </div>
          <span className="font-bold tracking-tight">NOTOPOS <span className="text-primary text-xs">AI</span></span>
        </div>

        <nav className="flex-1 p-4 space-y-1">
          {[
            { name: "Dashboard", icon: LayoutDashboard, active: true },
            { name: "POS / Sales", icon: ShoppingCart },
            { name: "Inventory", icon: Package },
            { name: "Customers", icon: Users },
            { name: "Analytics", icon: PieChart },
            { name: "Settings", icon: Settings },
          ].map((item) => (
            <a
              key={item.name}
              href="#"
              className={`flex items-center gap-3 px-4 py-3 rounded-xl transition-all font-medium text-sm ${item.active
                  ? "bg-primary text-primary-foreground shadow-md shadow-primary/20"
                  : "text-muted-foreground hover:bg-muted hover:text-foreground"
                }`}
            >
              <item.icon size={18} />
              {item.name}
            </a>
          ))}
        </nav>

        <div className="p-4 border-t">
          <div className="bg-primary/5 rounded-2xl p-4 border border-primary/10">
            <p className="text-xs font-bold text-primary mb-1 uppercase tracking-wider">AI Assistant</p>
            <p className="text-[10px] text-muted-foreground mb-3 leading-tight">
              Predicting 12% sales increase next week. Restock item &quot;Cappuccino&quot; now.
            </p>
            <button className="w-full bg-primary text-primary-foreground py-2 rounded-lg text-[10px] font-bold hover:opacity-90">
              View AI Report
            </button>
          </div>
        </div>
      </aside>

      {/* Main Content */}
      <main className="flex-1 flex flex-col overflow-hidden">
        {/* Header */}
        <header className="h-16 bg-background border-b px-8 flex items-center justify-between">
          <div className="relative w-96">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground" size={16} />
            <input
              type="text"
              placeholder="Search anything..."
              className="w-full bg-muted/50 border-none rounded-full pl-10 pr-4 py-2 text-sm focus:ring-1 ring-primary outline-none"
            />
          </div>

          <div className="flex items-center gap-4">
            <button className="bg-primary/10 text-primary p-2 rounded-full hover:bg-primary/20 transition-all relative">
              <Plus size={20} />
            </button>
            <button className="p-2 text-muted-foreground hover:bg-muted rounded-full relative">
              <Bell size={20} />
              <span className="absolute top-2 right-2 w-2 h-2 bg-destructive rounded-full border-2 border-background"></span>
            </button>
            <div className="h-8 w-px bg-border"></div>
            <div className="flex items-center gap-3 pl-2">
              <div className="text-right hidden md:block">
                <p className="text-xs font-bold">Sigit Wasis</p>
                <p className="text-[10px] text-muted-foreground">Store Owner</p>
              </div>
              <div className="w-9 h-9 rounded-full bg-gradient-to-tr from-primary to-blue-400"></div>
            </div>
          </div>
        </header>

        {/* Scrollable Area */}
        <div className="flex-1 overflow-y-auto p-8">
          {children}
        </div>
      </main>
    </div>
  );
}
