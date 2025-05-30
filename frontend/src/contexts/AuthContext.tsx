import { createContext, useContext, useState, useEffect } from 'react';
import { User as FirebaseUser, signInWithPopup, signOut, AuthError } from 'firebase/auth';
import { auth, googleProvider } from '../lib/firebase';
import toast from 'react-hot-toast';

interface AuthContextType {
  user: FirebaseUser | null;
  signInWithGoogle: () => Promise<void>;
  logout: () => Promise<void>;
  isAuthenticated: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<FirebaseUser | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const unsubscribe = auth.onAuthStateChanged((user) => {
      setUser(user);
      setLoading(false);
    });

    return () => unsubscribe();
  }, []);

  const signInWithGoogle = async () => {
    try {
      setLoading(true);
      const result = await signInWithPopup(auth, googleProvider);
      if (result.user) {
        toast.success('Successfully signed in!');
      }
    } catch (error) {
      const authError = error as AuthError;
      console.error('Google sign-in error:', authError);
      toast.error(`Failed to sign in: ${authError.message}`);
    } finally {
      setLoading(false);
    }
  };

  const logout = async () => {
    try {
      setLoading(true);
      await signOut(auth);
      toast.success('Successfully signed out!');
    } catch (error) {
      const authError = error as AuthError;
      console.error('Logout error:', authError);
      toast.error(`Failed to sign out: ${authError.message}`);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return <div>Loading authentication state...</div>;
  }

  return (
    <AuthContext.Provider value={{ user, signInWithGoogle, logout, isAuthenticated: !!user }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}