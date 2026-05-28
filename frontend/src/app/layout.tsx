import type { Metadata } from "next";
import "./globals.css";
import Providers from "@/lib/Providers";

export const metadata: Metadata = {
  title: "Hostel SaaS",
  description: "Hostel Management System",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <Providers>
        {children}
              </Providers>
      </body>
    </html>
  );
}
