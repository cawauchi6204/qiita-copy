import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'

import Header from "../components/Header"
import Footer from "../components/Footer"

const Home: NextPage = () => {
  return (
    <div>

      <main>
        <button className="bg-green-600 hover:bg-green-500 text-white rounded px-4 py-2">登録する</button>
      </main>
    </div>
  )
}

export default Home
