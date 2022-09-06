import Link from "next/link"
import { Mail, Bell, User } from 'react-feather'

type Props = {}

const Header: React.FC<Props> = ({ }) => {
  return (
    <header className="py-2 font-bold" style={{ "backgroundColor": "#2F3232" }}>
      <div className=" flex flex-col mx-8">
        <div className="flex justify-between">
          <Link href="" >
            <a className="text-green-500 font-normal text-4xl">
              Qiita
            </a>
          </Link>
          <div>
            <a className="text-white mr-4" href="">ベータ版フィードバック</a>
            <input className="w-72 h-10 text-xs p-2 font-normal" type="text" placeholder="記事を検索" />
            <Mail className="inline ml-4" color="white" size={24} />
            <Bell className="inline ml-4" color="white" size={24} />
            <User className="inline ml-4" color="white" size={24} />
            <button className="bg-green-500 text-white ml-4">投稿する</button>
          </div>
        </div>
        <ul className="flex text-gray-400 gap-6 mt-2">
          <li>
            <Link href="/">
              <a className="hover:text-white">
                ホーム
              </a>
            </Link>
          </li>
          <li>
            <Link href="/timeline">
              <a className="hover:text-white">
                タイムライン
              </a>
            </Link>
          </li>
          <li>
            <Link href="/trend">
              <a className="hover:text-white">
                トレンド
              </a>
            </Link>
          </li>
          <li>
            <Link href="/question">
              <a className="hover:text-white">
                質問
              </a>
            </Link>
          </li>
          <li>
            <Link href="/official-events">
              <a className="hover:text-white">
                公式イベント
              </a>
            </Link>
          </li>
          <li>
            <Link href="">
              <a className="hover:text-white">
                公式コラム
              </a>
            </Link>
          </li>
          <li>
            <Link href="/official-columns">
              <a className="hover:text-white">
                募集
              </a>
            </Link>
          </li>
          <li>
            <Link href="/organizations">
              <a className="hover:text-white">
                Organization
              </a>
            </Link>
          </li>
          <li>
            <Link href="/">
              <a className="hover:text-white">
                Qiita Blog
              </a>
            </Link>
          </li>
        </ul>
      </div>
    </header>
  )
}

export default Header
