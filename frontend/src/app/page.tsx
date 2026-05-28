import Link from 'next/link';

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24 bg-gray-50">
      <h1 className="text-4xl font-bold mb-8 text-gray-900">Hostel Management SaaS</h1>
      <div className="flex gap-4">
        <Link href="/login" className="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700">
          Login
        </Link>
        <Link href="/signup" className="px-6 py-2 bg-green-600 text-white rounded-md hover:bg-green-700">
          Sign Up
        </Link>
      </div>
    </main>
  );
}
