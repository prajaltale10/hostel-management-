'use client';
import { useQuery } from '@tanstack/react-query';
import api from '@/lib/api';
import Navbar from '@/components/Navbar';
import Sidebar from '@/components/Sidebar';

export default function Dashboard() {
  const { data: kpis, isLoading } = useQuery({
    queryKey: ['kpis'],
    queryFn: async () => {
      const res = await api.get('/analytics/kpis');
      return res.data;
    }
  });

  return (
    <div className="flex flex-col h-screen bg-gray-100">
      <Navbar />
      <div className="flex flex-1 overflow-hidden">
        <Sidebar />
        <main className="flex-1 p-6 overflow-y-auto">
          <h1 className="text-2xl font-bold mb-6 text-gray-800">Dashboard</h1>
          {isLoading ? (
            <p>Loading KPIs...</p>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div className="bg-white p-6 rounded-lg shadow border border-gray-200">
                <h3 className="text-gray-500 text-sm font-medium">Total Students</h3>
                <p className="text-3xl font-bold text-gray-800 mt-2">{kpis?.total_students || 0}</p>
              </div>
              <div className="bg-white p-6 rounded-lg shadow border border-gray-200">
                <h3 className="text-gray-500 text-sm font-medium">Occupied Beds</h3>
                <p className="text-3xl font-bold text-green-600 mt-2">{kpis?.occupied_beds || 0}</p>
              </div>
              <div className="bg-white p-6 rounded-lg shadow border border-gray-200">
                <h3 className="text-gray-500 text-sm font-medium">Vacant Beds</h3>
                <p className="text-3xl font-bold text-red-600 mt-2">{kpis?.vacant_beds || 0}</p>
              </div>
              <div className="bg-white p-6 rounded-lg shadow border border-gray-200">
                <h3 className="text-gray-500 text-sm font-medium">Total Revenue</h3>
                <p className="text-3xl font-bold text-blue-600 mt-2">₹{kpis?.total_revenue || 0}</p>
              </div>
              <div className="bg-white p-6 rounded-lg shadow border border-gray-200">
                <h3 className="text-gray-500 text-sm font-medium">Pending Dues</h3>
                <p className="text-3xl font-bold text-yellow-600 mt-2">₹{kpis?.pending_dues || 0}</p>
              </div>
            </div>
          )}
        </main>
      </div>
    </div>
  );
}
