"use client";

import React from "react";
import { 
  ArrowUpRight, 
  ArrowDownRight, 
  TrendingUp, 
  Package, 
  ShoppingBag, 
  Users,
  BrainCircuit,
  Lightbulb
} from "lucide-react";
import { motion } from "framer-motion";

export default function DashboardPage() {
  return (
    <div className="space-y-8 pb-12">
      <div className="flex justify-between items-end">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Business Overview</h1>
          <p className="text-muted-foreground mt-1">Here's what's happening with NOTOPOS AI today.</p>
        </div>
        <div className="flex gap-2">
            <button className="px-4 py-2 bg-background border rounded-xl text-sm font-medium hover:bg-muted transition-all">Download Report</button>
            <button className="px-4 py-2 bg-primary text-primary-foreground rounded-xl text-sm font-bold hover:opacity-90 transition-all">New Transaction</button>
        </div>
      </div>

      {/* Quick Stats */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        {[
          { label: "Total Revenue", val: "$12,450", trend: "+12.5%", icon: TrendingUp, pos: true },
          { label: "Orders Today", val: "156", trend: "+4.2%", icon: ShoppingBag, pos: true },
          { label: "Active Items", val: "1,240", trend: "-2.1%", icon: Package, pos: false },
          { label: "New Customers", val: "48", trend: "+18.3%", icon: Users, pos: true },
        ].map((stat, i) => (
          <motion.div 
            key={i}
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: i * 0.1 }}
            className="p-6 bg-background rounded-2xl border shadow-sm"
          >
            <div className="flex justify-between items-start mb-4">
              <div className="w-10 h-10 bg-primary/10 rounded-xl flex items-center justify-center text-primary">
                <stat.icon size={20} />
              </div>
              <div className={`flex items-center text-xs font-bold ${stat.pos ? 'text-green-500' : 'text-red-500'}`}>
                {stat.trend} {stat.pos ? <ArrowUpRight size={14} /> : <ArrowDownRight size={14} />}
              </div>
            </div>
            <p className="text-sm font-medium text-muted-foreground mb-1">{stat.label}</p>
            <p className="text-2xl font-extrabold tracking-tight">{stat.val}</p>
          </motion.div>
        ))}
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* Main Chart Area */}
        <div className="lg:col-span-2 p-8 bg-background rounded-3xl border shadow-sm min-h-[400px]">
          <div className="flex justify-between items-center mb-8">
            <h3 className="font-bold text-lg">Sales Performance</h3>
            <select className="bg-muted border-none rounded-lg px-3 py-1 text-xs outline-none">
              <option>Last 7 Days</option>
              <option>Last 30 Days</option>
            </select>
          </div>
          <div className="h-64 w-full bg-muted/20 rounded-2xl flex items-center justify-center relative overflow-hidden">
             {/* Chart Placeholder SVG */}
             <svg className="w-full h-full px-4" viewBox="0 0 400 100">
                <path 
                  d="M0,80 Q50,70 100,50 T200,60 T300,20 T400,30" 
                  fill="none" 
                  stroke="hsl(var(--primary))" 
                  strokeWidth="3"
                />
                <path 
                  d="M0,80 Q50,70 100,50 T200,60 T300,20 T400,30 V100 H0 Z" 
                  fill="url(#gradient)" 
                  className="opacity-10"
                />
                <defs>
                  <linearGradient id="gradient" x1="0" x2="0" y1="0" y2="1">
                    <stop offset="0%" stopColor="hsl(var(--primary))" />
                    <stop offset="100%" stopColor="transparent" />
                  </linearGradient>
                </defs>
             </svg>
             <div className="absolute bottom-4 left-0 w-full flex justify-between px-8 text-[10px] text-muted-foreground font-medium">
                <span>Mon</span><span>Tue</span><span>Wed</span><span>Thu</span><span>Fri</span><span>Sat</span><span>Sun</span>
             </div>
          </div>
        </div>

        {/* AI Recommendations */}
        <div className="space-y-6">
          <div className="p-6 bg-primary text-primary-foreground rounded-3xl shadow-lg shadow-primary/20">
            <div className="flex items-center gap-2 mb-4">
              <BrainCircuit size={20} />
              <h3 className="font-bold">AI Business Insights</h3>
            </div>
            <div className="space-y-4">
              <div className="bg-white/10 p-4 rounded-2xl border border-white/10">
                <p className="text-xs font-medium mb-1 opacity-80">Sales Forecast</p>
                <p className="text-sm font-bold">Revenue expected to rise by 15% this weekend due to local event.</p>
              </div>
              <div className="bg-white/10 p-4 rounded-2xl border border-white/10">
                <p className="text-xs font-medium mb-1 opacity-80">Inventory Alert</p>
                <p className="text-sm font-bold">Stock of 'Milk' will run out in 2 days. Reorder now to avoid loss.</p>
              </div>
            </div>
          </div>

          <div className="p-6 bg-background rounded-3xl border shadow-sm">
             <div className="flex items-center gap-2 mb-4 text-orange-500">
              <Lightbulb size={20} />
              <h3 className="font-bold text-foreground">Smart Suggestions</h3>
            </div>
            <ul className="space-y-4 text-sm">
              <li className="flex items-start gap-3">
                <div className="w-2 h-2 rounded-full bg-primary mt-1.5 shrink-0"></div>
                <p className="text-muted-foreground">Bundle <span className="font-bold text-foreground">Croissant</span> with <span className="font-bold text-foreground">Latte</span> for a 5% increase in basket size.</p>
              </li>
              <li className="flex items-start gap-3">
                <div className="w-2 h-2 rounded-full bg-primary mt-1.5 shrink-0"></div>
                <p className="text-muted-foreground">Traffic peaks at <span className="font-bold text-foreground">4 PM</span>. Add 1 more staff to shift.</p>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}
