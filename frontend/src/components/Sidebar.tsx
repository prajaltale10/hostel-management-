import Link from 'next/link';

export default function Sidebar() {
  const links = [
    { name: 'Dashboard', href: '/dashboard' },
    { name: 'Hostels', href: '/hostels' },
    { name: 'Students', href: '/students' },
  ];

  return (
    <aside className="w-64 bg-gray-800 text-white min-h-screen p-4">
      <nav className="space-y-2">
        {links.map((link) => (
          <Link key={link.name} href={link.href} className="block px-4 py-2 rounded hover:bg-gray-700">
            {link.name}
          </Link>
        ))}
      </nav>
    </aside>
  );
}
