'use client';
import { useAuthStore } from '@/store/useStore';
import { useRouter } from 'next/navigation';

export default function Navbar() {
  const { user, logout } = useAuthStore();
  const router = useRouter();

  const handleLogout = () => {
    logout();
    router.push('/login');
  };

  return (
    <nav className="bg-white border-b border-gray-200 px-4 py-3 flex justify-between items-center">
      <div className="text-xl font-bold text-gray-800">Hostel SaaS</div>
      <div className="flex items-center gap-4">
        {user && <span className="text-sm text-gray-600">{user.name} ({user.role})</span>}
        <button onClick={handleLogout} className="text-sm text-red-600 hover:text-red-800">Logout</button>
      </div>
    </nav>
  );
}
