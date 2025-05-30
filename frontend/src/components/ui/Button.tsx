import React from 'react';
import { cn } from '../../utils/cn';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'outline' | 'ghost';
  size?: 'sm' | 'md' | 'lg';
  fullWidth?: boolean;
  children: React.ReactNode;
}

const Button: React.FC<ButtonProps> = ({
  variant = 'primary',
  size = 'md',
  fullWidth = false,
  className,
  children,
  ...props
}) => {
  return (
    <button
      className={cn(
        'inline-flex items-center justify-center rounded-md font-medium transition-all duration-300 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-purple-600 disabled:opacity-50 shadow-lg',
        {
          'bg-purple-gradient text-white hover:opacity-90 hover:scale-105': variant === 'primary',
          'bg-gradient-to-r from-gray-800 to-gray-900 text-white hover:opacity-90': variant === 'secondary',
          'border border-purple-500 bg-transparent hover:bg-purple-gradient hover:border-transparent hover:text-white': variant === 'outline',
          'bg-transparent hover:bg-gray-900 hover:text-white': variant === 'ghost',
          'h-9 px-4 text-sm': size === 'sm',
          'h-11 px-6 text-base': size === 'md',
          'h-14 px-8 text-lg': size === 'lg',
          'w-full': fullWidth,
        },
        className
      )}
      {...props}
    >
      {children}
    </button>
  );
};

export default Button;