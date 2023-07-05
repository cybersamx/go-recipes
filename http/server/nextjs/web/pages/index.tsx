import Head from 'next/head';
import { useState } from "react";

export default function Home() {
  const [name, setName] = useState('');

  async function refresh() {
    try {
      const res = await fetch(`http://localhost:3000/api`);
      const name = await res.text();
      setName(name);
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <>
      <Head>
        <title>Go Recipes</title>
      </Head>
      <div>
        <h1>Next.js served from Go server</h1>
        <p>{name}</p>
        <button onClick={refresh}>Refresh</button>
      </div>
    </>
  );
}
