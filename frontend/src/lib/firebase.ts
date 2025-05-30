import { initializeApp } from 'firebase/app';
import { getAuth, GoogleAuthProvider } from 'firebase/auth';

const firebaseConfig = {
  apiKey: "AIzaSyAUjQq-a2yfzZgH2MCjZbLcRKQPLWq6jI0",
  authDomain: "car-rental-2ea40.firebaseapp.com",
  projectId: "car-rental-2ea40",
  storageBucket: "car-rental-2ea40.firebasestorage.app",
  messagingSenderId: "400717904530",
  appId: "1:400717904530:web:7d722d615716538c92044b",
  measurementId: "G-ZPP8SKNVSW"
};
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
export const googleProvider = new GoogleAuthProvider();