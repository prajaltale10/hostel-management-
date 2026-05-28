'use client';
import { useState } from 'react';
import { useAuthStore } from '@/store/useStore';
import api from '@/lib/api';
import { useRouter } from 'next/navigation';
import Link from 'next/link';

export default function Signup() {
  const [orgName, setOrgName] = useState('');
  const [userName, setUserName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const setAuth = useAuthStore((state) => state.setAuth);
  const router = useRouter();

  const handleSignup = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await api.post('/auth/signup', { org_name: orgName, user_name: userName, email, password });
      setAuth(res.data.user, res.data.token);
      router.push('/dashboard');
    } catch (err: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
      setError(err.response?.data?.error || 'Signup failed');
    }
  };

  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-50">
      <div className="w-full max-w-md p-8 bg-white rounded-lg shadow">
        <h2 className="text-2xl font-bold mb-6 text-center text-black">Sign Up</h2>
        {error && <p className="text-red-500 mb-4 text-center">{error}</p>}
        <form onSubmit={handleSignup} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">Organization Name</label>
            <input type="text" value={orgName} onChange={(e) => setOrgName(e.target.value)} required className="mt-1 block w-full p-2 border border-gray-300 rounded text-black" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Your Name</label>
            <input type="text" value={userName} onChange={(e) => setUserName(e.target.value)} required className="mt-1 block w-full p-2 border border-gray-300 rounded text-black" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Email</label>
            <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required className="mt-1 block w-full p-2 border border-gray-300 rounded text-black" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Password</label>
            <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required className="mt-1 block w-full p-2 border border-gray-300 rounded text-black" />
          </div>
          <button type="submit" className="w-full py-2 px-4 bg-green-600 text-white rounded hover:bg-green-700">Sign Up</button>
        </form>
        <p className="mt-4 text-center text-sm text-gray-600">
          Already have an account? <Link href="/login" className="text-blue-600">Login</Link>
        </p>
      </div>
    </div>
  );
}
