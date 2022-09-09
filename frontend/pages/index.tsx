import type { GetServerSideProps, NextPage } from 'next'
import Layout from "../components/Layout"
import CardList from "../components/CardList"
import Ranking from "../components/Ranking"
import axios from 'axios'
import { Post } from '../types/Post'

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

type Props = {
  posts: Post[]
}

const setCookie = async () => {
  await axios.get('/api/set_cookie-by-front')
}
const getCookie = async () => {
  await axios.get('/api/get_cookie-by-front')
}
const deleteCookie = async () => {
  await axios.get('/api/delete_cookie-by-front')
}

const Home: NextPage<Props> = ({ posts }) => {
  return (
    <Layout>
      <div className="flex gap-8">
        <div className="w-1/6">
          <Ranking rankingTitle="タグランキング" rankingUnit="users" contents={dummyRanking} />
        </div>
        <div className="w-2/3">
          <button onClick={() => setCookie()} className="bg-blue-600 hover:bg-red-500 text-white rounded px-4 py-2">クッキーを設定する</button>
          <button onClick={() => getCookie()} className="bg-blue-600 hover:bg-red-500 text-white rounded px-4 py-2">クッキーを取得する</button>
          <button onClick={() => deleteCookie()} className="bg-blue-600 hover:bg-red-500 text-white rounded px-4 py-2">クッキーを削除する</button>
          <CardList posts={posts} />
        </div>
        <div className="w-1/6">
          <Ranking rankingTitle="タグランキング" rankingUnit="users" contents={dummyRanking} />
        </div>
      </div>
    </Layout>
  )
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const res = await axios.get(`http://localhost/posts`)
  return { props: { posts: res.data } }
}

export default Home
