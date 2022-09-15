import Link from "next/link"
import { Mail, Bell, User } from 'react-feather'
import {
  Menu,
  MenuItem,
  MenuButton,
} from '@szhsin/react-menu';
import { useRecoilValue } from "recoil";
import { userState } from "../contexts/UserContext";
import Icon from "./Icon";

type Props = {
  showNavBars?: boolean
}

const Header: React.FC<Props> = ({ showNavBars = true }) => {
  const user = useRecoilValue(userState);
  return (
    <header className="py-2 font-bold" style={{ "backgroundColor": "#2F3232" }}>
      <div className=" flex flex-col mx-8">
        <div className="flex justify-between">
          <Link href="/" >
            <a className="text-green-500 font-normal text-4xl">
              Qiita
            </a>
          </Link>
          <div className="flex items-center">
            <a className="text-white text-xs mr-4 hover:underline" href="">ベータ版フィードバック</a>
            <input className="w-72 h-10 text-xs p-2 font-normal" type="text" placeholder="記事を検索" />
            {user && user.id !== "" ? (
              <>
                <Menu className="inline" menuButton={<MenuButton><Mail className="inline ml-4" color="white" size={24} /></MenuButton>}>
                  <Link href="/mypage">
                    <a>
                      <MenuItem className="text-xs font-normal p-4 w-40 bg-white hover:bg-gray-300">
                        hogehogeさんから連絡がありました
                      </MenuItem>
                    </a>
                  </Link>
                </Menu>
                <Menu className="inline" menuButton={<MenuButton><Bell className="inline ml-4" color="white" size={24} /></MenuButton>}>
                  <Link href="/mypage">
                    <a>
                      <MenuItem className="text-xs font-normal p-4 w-40 bg-white hover:bg-gray-300">
                        hogehogeさんからfollowされました
                      </MenuItem>
                    </a>
                  </Link>
                </Menu>
                <Menu className="inline" menuButton={<MenuButton><div className="ml-4"><Icon src={user.imgUrl} width={24} height={24} /></div></MenuButton>}>
                  <Link href={`/${user.id}`}>
                    <a>
                      <MenuItem className="text-xs font-normal p-4 w-40 bg-white hover:bg-gray-300">
                        マイページ
                      </MenuItem>
                    </a>
                  </Link>
                  <Link href="/stock">
                    <a>
                      <MenuItem className="text-xs font-normal p-4 w-40 bg-white hover:bg-gray-300">
                        ストック
                      </MenuItem>
                    </a>
                  </Link>
                  <Link href="/drafts/new">
                    <a>
                      <MenuItem className="text-xs font-normal p-4 w-40 bg-white hover:bg-gray-300">
                        投稿
                      </MenuItem>
                    </a>
                  </Link>
                  <Link href="/drafts">
                    <a>
                      <MenuItem className="text-xs font-normal p-4 w-40 bg-white hover:bg-gray-300">
                        下書き
                      </MenuItem>
                    </a>
                  </Link>
                  <Link href="/settings/profile">
                    <a>
                      <MenuItem className="text-xs font-normal p-4 w-40 bg-white hover:bg-gray-300">
                        設定
                      </MenuItem>
                    </a>
                  </Link>
                </Menu>
                <Link href="/drafts/new">
                  <button className="bg-green-500 text-white ml-4">投稿する</button>
                </Link>
              </>
            )
              :
              (
                <>
                  <Link href="/login">
                    <a className="text-xs text-white ml-2">
                      ログイン
                    </a>
                  </Link>
                  <Link href="/signup">
                    <button className=" text-xs border-white border-2 border-solid text-white p-2 ml-2">新規登録</button>
                  </Link>
                </>
              )
            }
          </div>
        </div>
        {showNavBars && (
          <ul className="flex text-gray-400 text-medium gap-6 mt-2 mx-4">
            <li>
              <Link href="/">
                <a className="hover:text-white">
                  ホーム
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
        )}
      </div>
    </header>
  )
}

export default Header
