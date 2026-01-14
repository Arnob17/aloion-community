import Image from "next/image";
import Link from "next/link";
export default function Navbar() {
  const items = [
    {
      title: "Home",
      url: "/",
    },
    {
      title: "Organizations",
      url: "/org",
    },
  ];

  return (
    <div className="p-5 w-full bg-black h-20 flex justify-between items-center sticky top-0">
      <div className="w-30 flex justify-center">
        <Link href={"/"}>
          <Image
            src={"/logo/aloionLogo.jpg"}
            height={68}
            width={68}
            alt="aloion_logo"
            className="rounded-full"
          />
        </Link>
      </div>
      <div className="w-auto flex justify-between">
        <ul className="flex gap-4">
          {items.map((item) => (
            <li
              key={item.title}
              className="text-white font-bold hover:bg-white hover:text-black hover:rounded-[5px] cursor-pointer"
            >
              <Link href={item.url}>{item.title}</Link>
            </li>
          ))}
        </ul>
      </div>
      <div className="flex items-center gap-3 cursor-pointer">
        <div className="text-white font-bold"> Arnob </div>
        <div className="rounded-full h-10 w-10 bg-white text-black flex justify-center items-center">
          U
        </div>
      </div>
    </div>
  );
}
