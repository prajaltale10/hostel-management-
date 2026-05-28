'use client';
import { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import api from '@/lib/api';
import Navbar from '@/components/Navbar';
import Sidebar from '@/components/Sidebar';
import Table from '@/components/Table';
import Modal from '@/components/Modal';

export default function Hostels() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [name, setName] = useState('');
  const [address, setAddress] = useState('');
  const queryClient = useQueryClient();

  const { data: hostels = [], isLoading } = useQuery({
    queryKey: ['hostels'],
    queryFn: async () => {
      const res = await api.get('/hostels/');
      return res.data;
    }
  });

  const mutation = useMutation({
    mutationFn: (newHostel: any /* eslint-disable-line @typescript-eslint/no-explicit-any */) => api.post('/hostels/', newHostel),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['hostels'] });
      setIsModalOpen(false);
      setName('');
      setAddress('');
    }
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    mutation.mutate({ name, address });
  };

  const columns = [
    { header: 'ID', accessor: 'id' },
    { header: 'Name', accessor: 'name' },
    { header: 'Address', accessor: 'address' },
  ];

  return (
    <div className="flex flex-col h-screen bg-gray-100">
      <Navbar />
      <div className="flex flex-1 overflow-hidden">
        <Sidebar />
        <main className="flex-1 p-6 overflow-y-auto">
          <div className="flex justify-between items-center mb-6">
            <h1 className="text-2xl font-bold text-gray-800">Hostels</h1>
            <button onClick={() => setIsModalOpen(true)} className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
              Add Hostel
            </button>
          </div>

          {isLoading ? <p>Loading...</p> : <Table columns={columns} data={hostels} />}

          <Modal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)} title="Add New Hostel">
            <form onSubmit={handleSubmit} className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700">Name</label>
                <input type="text" value={name} onChange={(e) => setName(e.target.value)} required className="mt-1 block w-full p-2 border border-gray-300 rounded text-black" />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-700">Address</label>
                <input type="text" value={address} onChange={(e) => setAddress(e.target.value)} className="mt-1 block w-full p-2 border border-gray-300 rounded text-black" />
              </div>
              <button type="submit" disabled={mutation.isPending} className="w-full py-2 px-4 bg-blue-600 text-white rounded hover:bg-blue-700">
                {mutation.isPending ? 'Saving...' : 'Save'}
              </button>
            </form>
          </Modal>
        </main>
      </div>
    </div>
  );
}
