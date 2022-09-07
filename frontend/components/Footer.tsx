import Link from "next/link"

type Props = {}

const Footer: React.FC<Props> = ({ }) => {
  return (
    <footer className="py-12 text-white" style={{ "backgroundColor": "#2F3232" }}>
      <div className=" flex mx-8 justify-between">
        <div className="flex flex-col justify-between">
          <div>
            <Link href="/">
              <a className="text-4xl">
                Qiita
              </a>
            </Link>
            <p className="mt-6">
              How developers code is here.
            </p>
          </div>
          <small>&copy; 2011-2022Qiita Inc.</small>
        </div>
        <div className="flex gap-12">
          <ul className="flex flex-col gap-4">
            <li className="font-bold">
              ガイドとヘルプ
            </li>
            <li>
              <Link href="">
                <a>
                  about
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  利用規約
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>プライバシーポリシー</a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>ガイドライン</a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>デザインガイドライン</a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>ご意見</a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>ヘルプ</a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>広告掲載</a>
              </Link>
            </li>
          </ul>
          <ul className="flex flex-col gap-4">
            <li className="font-bold">
              コンテンツ
            </li>
            <li>
              <Link href="">
                <a>
                  リリースノート
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  公式イベント
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  公式コラム
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  募集
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  アドベントカレンダー
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  Qiita 表彰プログラム
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>API</a>
              </Link>
            </li>
          </ul>
          <ul className="flex flex-col gap-4">
            <li className="font-bold">
              SNS
            </li>
            <li>
              <Link href="">
                <a>
                  Qiita（キータ）公式
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  Qiita マイルストーン
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  Qiita 人気の投稿
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  Qiita（キータ）公式
                </a>
              </Link>
            </li>
          </ul>
          <ul className="flex flex-col gap-4">
            <li className="font-bold">
              Qiita 関連サービス
            </li>
            <li>
              <Link href="">
                <a>
                  Qiita Team
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  Qiita Jobs
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  Qiita Zine
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  Qiita 公式ショップ
                </a>
              </Link>
            </li>
          </ul>
          <ul className="flex flex-col gap-4">
            <li className="font-bold">
              運営
            </li>
            <li>
              <Link href="">
                <a>
                  運営会社
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  採用情報
                </a>
              </Link>
            </li>
            <li>
              <Link href="">
                <a>
                  Qiita Blog
                </a>
              </Link>
            </li>
          </ul>
        </div>
      </div>
    </footer>
  )
}

export default Footer
