'use client';
import { useState } from 'react';
import { useAuthStore } from '@/store/useStore';
import api from '@/lib/api';
import { useRouter } from 'next/navigation';
import Link from 'next/link';

export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const setAuth = useAuthStore((state) => state.setAuth);
  const router = useRouter();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await api.post('/auth/login', { email, password });
      setAuth(res.data.user, res.data.token);
      router.push('/dashboard');
    } catch (err: any) { // eslint-disable-line @typescript-eslint/no-explicit-any
      setError(err.response?.data?.error || 'Login failed');
    }
  };

  return (
    <div className="flex min-h-screen items-center justify-center bg-gray-50">
      <div className="w-full max-w-md p-8 bg-white rounded-lg shadow">
        <h2 className="text-2xl font-bold mb-6 text-center text-black">Login</h2>
        {error && <p className="text-red-500 mb-4 text-center">{error}</p>}
        <form onSubmit={handleLogin} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">Email</label>
            <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required className="mt-1 block w-full p-2 border border-gray-300 rounded text-black" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Password</label>
            <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required className="mt-1 block w-full p-2 border border-gray-300 rounded text-black" />
          </div>
          <button type="submit" className="w-full py-2 px-4 bg-blue-600 text-white rounded hover:bg-blue-700">Login</button>
        </form>
        <p className="mt-4 text-center text-sm text-gray-600">
          Don&apos;t have an account? <Link href="/signup" className="text-blue-600">Sign up</Link>
        </p>
      </div>
    </div>
  );
}
