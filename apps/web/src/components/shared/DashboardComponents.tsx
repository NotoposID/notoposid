import React from "react";
import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

interface CardProps {
  children: React.ReactNode;
  className?: string;
  onClick?: () => void;
}

export const Card = ({ children, className, onClick }: CardProps) => {
  return (
    <div 
      onClick={onClick}
      className={cn(
        "bg-background border rounded-3xl p-6 shadow-sm hover:shadow-md transition-all duration-300",
        onClick && "cursor-pointer active:scale-[0.98]",
        className
      )}
    >
      {children}
    </div>
  );
};

interface StatProps {
  label: string;
  value: string | number;
  trend?: string;
  trendUp?: boolean;
  icon: React.ReactNode;
  className?: string;
}

export const StatItem = ({ label, value, trend, trendUp, icon, className }: StatProps) => {
  return (
    <Card className={cn("flex flex-col", className)}>
      <div className="flex justify-between items-start mb-4">
        <div className="w-12 h-12 bg-primary/10 rounded-2xl flex items-center justify-center text-primary">
          {icon}
        </div>
        {trend && (
          <div className={cn(
            "text-xs font-bold px-2 py-1 rounded-full",
            trendUp ? "text-green-600 bg-green-50" : "text-red-600 bg-red-50"
          )}>
            {trend}
          </div>
        )}
      </div>
      <div>
        <p className="text-sm font-medium text-muted-foreground mb-1">{label}</p>
        <h3 className="text-2xl font-bold tracking-tight">{value}</h3>
      </div>
    </Card>
  );
};
