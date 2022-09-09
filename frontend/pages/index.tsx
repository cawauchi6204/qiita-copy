import type { NextPage } from 'next'
import Layout from "../components/Layout"
import CardList from "../components/CardList"
import Ranking from "../components/Ranking"

const dummyCards = [
  {
    id: "hogehoge",
    title: "title1",
    createdAt: "1月"
  },
  {
    id: "fugafuga",
    title: "title2",
    createdAt: "2月"
  },
  {
    id: "piyopiyo",
    title: "title3",
    createdAt: "3月"
  },
]

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

const Home: NextPage = () => {
  return (
    <Layout>
      <div className="flex gap-8">
        <div className="w-1/6">
          <Ranking rankingTitle="タグランキング" rankingUnit="users" contents={dummyRanking} />
        </div>
        <div className="w-2/3">
          <CardList posts={dummyCards} />
        </div>
        <div className="w-1/6">
          <Ranking rankingTitle="タグランキング" rankingUnit="users" contents={dummyRanking} />
        </div>
      </div>
    </Layout>
  )
}

export default Home
