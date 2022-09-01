import React from "react"

type Props = {}

const SignUp: React.FC<Props> = ({ }) => {
  return (
    <>
      <h2>
        Qiitaへようこそ
      </h2>
      <p>新規登録(無料)して利用を開始しましょう。</p>
      <div>
        <input className="w-full" type="text" placeholder="ユーザー名" /><br />
        <input className="w-full" type="email" placeholder="メールアドレス" /><br />
        <input className="w-full" type="password" placeholder="パスワード" /><br />
        <input type="checkbox" />利用規約に同意する<br />
        <input type="checkbox" />プライバシーポリシーに同意する<br />
        <button className="bg-green-600 hover:bg-green-500 text-white rounded px-4 py-2">登録する</button>
      </div>
    </>
  )
}

export default SignUp
