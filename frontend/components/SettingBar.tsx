import Link from "next/link"

type Props = {}

const SettingBar: React.FC<Props> = ({ }) => {
  return (
    <ul className="bg-white p-4">
      <li className="text-xs p-2"><Link href="/settings/profile"><a className="font-bold">公開用プロフィール</a></Link></li><hr />
      <li className="text-xs p-2"><Link href="/settings/account"><a>アカウント</a></Link></li><hr />
      <li className="text-xs p-2"><Link href="/settings/password"><a>パスワード</a></Link></li><hr />
      <li className="text-xs p-2"><Link href="/settings/notifications"><a>メールアドレスと通知</a></Link></li><hr />
      <li className="text-xs p-2"><Link href="/settings/security"><a>セキュリティ</a></Link></li><hr />
      <li className="text-xs p-2"><Link href="/settings/two_factor_authentication/intro"><a>二段階認証</a></Link></li><hr />
      <li className="text-xs p-2"><Link href="/settings/applications"><a>アプリケーション</a></Link></li><hr />
      <li className="text-xs p-2"><Link href="/settings/mutes"><a>ミュート設定</a></Link></li><hr />
      <li className="text-xs p-2"><Link href="/settings/beta-release"><a>ベータ版切り替え</a></Link></li><hr />
      <li className="text-xs p-2"><Link href="/settings/analytics"><a>Googleアナリティクス設定</a></Link></li><hr />
      <li className="text-xs p-2"><Link href="/settings/organizations"><a>Organization設定</a></Link></li>
    </ul>
  )
}

export default SettingBar
