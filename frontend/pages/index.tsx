import type { NextPage } from 'next'
import Layout from "../components/Layout"
import CardList from "../components/CardList"
import Ranking from "../components/Ranking"
import axios from "axios"

const dummyRanking = [
  {
    img: "",
    name: "Python",
    count: 127,
  },
  {
    img: "",
    name: "Python",
    count: 127,
  },
  {
    img: "",
    name: "Python",
    count: 127,
  },
  {
    img: "",
    name: "Python",
    count: 127,
  },
  {
    img: "",
    name: "Python",
    count: 127,
  },
]

const setCookie = async () => {
  await axios.get('http://localhost/set_cookie-by-front'),
    { withCredentials: true }
}
const getCookie = async () => {
  await axios.get('http://localhost/get_cookie-by-front'),
    { withCredentials: true }
}
const deleteCookie = async () => {
  await axios.get('http://localhost/delete_cookie-by-front'),
    { withCredentials: true }
}

const Home: NextPage = () => {
  return (
    <Layout>
      <div className="flex gap-8">
        <div className="w-1/6">
          <Ranking rankingTitle="タグランキング" rankingUnit="users" contents={dummyRanking} />
        </div>
        <div className="w-2/3">
          <button onClick={() => setCookie()} className="bg-blue-600 hover:bg-red-500 text-white rounded px-4 py-2">クッキーを設定する</button>
          <button onClick={() => getCookie()} className="bg-blue-600 hover:bg-red-500 text-white rounded px-4 py-2">クッキーを取得する</button>
          <button onClick={() => deleteCookie()} className="bg-blue-600 hover:bg-red-500 text-white rounded px-4 py-2">クッキーを削除する</button>Ï
          <CardList />
        </div>
        <div className="w-1/6">
          <Ranking rankingTitle="タグランキング" rankingUnit="users" contents={dummyRanking} />
        </div>
      </div>
    </Layout>
  )
}

export default Home
