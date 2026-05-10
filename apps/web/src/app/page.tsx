"use client";

import React from "react";
import { motion } from "framer-motion";
import { 
  Bot, 
  BarChart3, 
  ShoppingBag, 
  Users, 
  ShieldCheck, 
  Zap,
  ArrowRight,
  TrendingUp,
  BrainCircuit
} from "lucide-react";

export default function LandingPage() {
  return (
    <div className="min-h-screen bg-background text-foreground selection:bg-primary/20">
      {/* Navigation */}
      <nav className="fixed top-0 w-full z-50 glass border-b">
        <div className="container mx-auto px-6 py-4 flex justify-between items-center">
          <div className="flex items-center gap-2">
            <div className="w-10 h-10 bg-primary rounded-xl flex items-center justify-center text-primary-foreground font-bold text-xl">
              N
            </div>
            <span className="text-2xl font-bold tracking-tight">NOTOPOS <span className="text-primary">AI</span></span>
          </div>
          <div className="hidden md:flex items-center gap-8 text-sm font-medium">
            <a href="#features" className="hover:text-primary transition-colors">Features</a>
            <a href="#ai" className="hover:text-primary transition-colors">AI Intelligence</a>
            <a href="#pricing" className="hover:text-primary transition-colors">Pricing</a>
            <button className="bg-primary text-primary-foreground px-6 py-2.5 rounded-full hover:opacity-90 transition-all font-semibold">
              Get Started
            </button>
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="relative pt-40 pb-20 overflow-hidden hero-gradient">
        <div className="container mx-auto px-6 text-center">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.6 }}
          >
            <div className="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-primary/10 text-primary text-sm font-bold mb-8 border border-primary/20">
              <Zap size={16} /> Enterprise-Grade AI POS for Modern Business
            </div>
            <h1 className="text-5xl md:text-7xl font-extrabold tracking-tight mb-8 leading-[1.1]">
              Scale Your Business with <br />
              <span className="text-primary">AI-Powered</span> Intelligence
            </h1>
            <p className="max-w-2xl mx-auto text-muted-foreground text-lg mb-10 leading-relaxed">
              The only POS system that doesn't just record sales—it predicts them. 
              Integrated with cutting-edge AI to help UMKM, Cafes, and Retailers grow faster.
            </p>
            <div className="flex flex-col md:flex-row gap-4 justify-center items-center">
              <button className="bg-primary text-primary-foreground px-8 py-4 rounded-xl text-lg font-bold hover:scale-105 transition-all flex items-center gap-2 shadow-lg shadow-primary/25">
                Start Free Trial <ArrowRight size={20} />
              </button>
              <button className="px-8 py-4 rounded-xl text-lg font-bold border hover:bg-muted transition-all">
                Book a Demo
              </button>
            </div>
          </motion.div>

          {/* Dashboard Preview Mockup */}
          <motion.div 
            initial={{ opacity: 0, y: 40 }}
            whileInView={{ opacity: 1, y: 0 }}
            viewport={{ once: true }}
            transition={{ duration: 0.8, delay: 0.2 }}
            className="mt-20 relative"
          >
            <div className="max-w-5xl mx-auto p-4 glass rounded-2xl shadow-2xl border border-white/10">
              <div className="aspect-video bg-muted/30 rounded-lg flex items-center justify-center overflow-hidden">
                <img 
                   src="https://images.unsplash.com/photo-1551288049-bebda4e38f71?q=80&w=2000&auto=format&fit=crop" 
                   alt="Dashboard Preview"
                   className="w-full h-full object-cover opacity-80"
                />
              </div>
            </div>
            {/* Floating Stats */}
            <div className="absolute -bottom-10 left-1/2 -translate-x-1/2 flex gap-4 w-full justify-center px-6">
                {[
                  { label: "Sales Growth", val: "+24%", icon: TrendingUp, color: "text-green-500" },
                  { label: "AI Insights", val: "Real-time", icon: BrainCircuit, color: "text-primary" },
                  { label: "Uptime", val: "99.9%", icon: ShieldCheck, color: "text-blue-500" }
                ].map((stat, i) => (
                  <motion.div 
                    key={i}
                    whileHover={{ y: -5 }}
                    className="glass p-6 rounded-2xl min-w-[200px] border border-white/10 shadow-xl"
                  >
                    <div className="flex items-center justify-between mb-2">
                      <stat.icon className={stat.color} size={24} />
                      <span className={`font-bold ${stat.color}`}>{stat.val}</span>
                    </div>
                    <div className="text-sm font-medium text-muted-foreground">{stat.label}</div>
                  </motion.div>
                ))}
            </div>
          </motion.div>
        </div>
      </section>

      {/* AI Features Section */}
      <section id="ai" className="py-32 bg-muted/30">
        <div className="container mx-auto px-6">
          <div className="text-center mb-20">
            <h2 className="text-4xl font-bold mb-4">Intelligent Business Core</h2>
            <p className="text-muted-foreground">More than just a cash register—a business brain.</p>
          </div>
          
          <div className="grid md:grid-cols-3 gap-8">
            {[
              {
                title: "Predictive Analytics",
                desc: "AI models forecast your sales for the next 7-30 days with 90%+ accuracy.",
                icon: BarChart3
              },
              {
                title: "Smart Inventory",
                desc: "Automatic restock alerts and waste reduction powered by sales pattern analysis.",
                icon: ShoppingBag
              },
              {
                title: "AI Business Assistant",
                desc: "Chat with your POS to get instant reports and business advice.",
                icon: Bot
              }
            ].map((feature, i) => (
              <div key={i} className="p-8 rounded-3xl bg-background border hover:border-primary/50 transition-all group">
                <div className="w-14 h-14 bg-primary/10 rounded-2xl flex items-center justify-center mb-6 group-hover:bg-primary group-hover:text-primary-foreground transition-all">
                  <feature.icon size={28} />
                </div>
                <h3 className="text-xl font-bold mb-4">{feature.title}</h3>
                <p className="text-muted-foreground leading-relaxed">
                  {feature.desc}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="py-20 border-t">
        <div className="container mx-auto px-6">
          <div className="grid md:grid-cols-4 gap-12 mb-16">
            <div className="col-span-1 md:col-span-1">
              <div className="flex items-center gap-2 mb-6">
                <div className="w-8 h-8 bg-primary rounded-lg flex items-center justify-center text-primary-foreground font-bold">
                  N
                </div>
                <span className="text-xl font-bold tracking-tight">NOTOPOS</span>
              </div>
              <p className="text-sm text-muted-foreground leading-relaxed">
                Empowering UMKM with enterprise-grade AI intelligence. Scale faster, manage smarter.
              </p>
            </div>
            <div>
              <h4 className="font-bold mb-6 text-sm uppercase tracking-widest text-primary">Product</h4>
              <ul className="space-y-4 text-sm text-muted-foreground">
                <li><a href="#" className="hover:text-primary transition-colors">POS Transaction</a></li>
                <li><a href="#" className="hover:text-primary transition-colors">Inventory Management</a></li>
                <li><a href="#" className="hover:text-primary transition-colors">AI Insights</a></li>
              </ul>
            </div>
            <div>
              <h4 className="font-bold mb-6 text-sm uppercase tracking-widest text-primary">Support</h4>
              <ul className="space-y-4 text-sm text-muted-foreground">
                <li><a href="#" className="hover:text-primary transition-colors">Documentation</a></li>
                <li><a href="#" className="hover:text-primary transition-colors">API Reference</a></li>
                <li><a href="#" className="hover:text-primary transition-colors">Community</a></li>
              </ul>
            </div>
            <div>
              <h4 className="font-bold mb-6 text-sm uppercase tracking-widest text-primary">Company</h4>
              <ul className="space-y-4 text-sm text-muted-foreground">
                <li><a href="#" className="hover:text-primary transition-colors">About Us</a></li>
                <li><a href="#" className="hover:text-primary transition-colors">Contact</a></li>
                <li><a href="#" className="hover:text-primary transition-colors">Privacy Policy</a></li>
              </ul>
            </div>
          </div>
          <div className="flex flex-col md:flex-row justify-between items-center pt-8 border-t text-sm text-muted-foreground gap-4">
            <p>© 2024 NOTOPOS AI. All rights reserved.</p>
            <div className="flex gap-8">
              <a href="#" className="hover:text-primary">Twitter</a>
              <a href="#" className="hover:text-primary">LinkedIn</a>
              <a href="#" className="hover:text-primary">GitHub</a>
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}
