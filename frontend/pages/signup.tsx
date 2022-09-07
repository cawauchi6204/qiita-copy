import { useState } from "react"
import axios from "axios"
import Layout from "../components/Layout"

// TODO: 要リファクタ(汚すぎる)
const SignUp = ({ }) => {
  const [name, setName] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const handleSubmit = async () => {
    await axios.post('http://localhost/signup', {
      name, email, password
    })
  }
  const getUser = async () => {
    await axios.get('http://localhost/users')
  }
  const login = async () => {
    await axios.post('http://localhost/login', {
      email, password
    })
  }
  const logout = async () => {
    await axios.post('http://localhost/logout')
  }
  const mypage = async () => {
    await axios.get('http://localhost/mypage')
  }
  return (
    <Layout>
      <h2>
        Qiitaへようこそ
      </h2>
      <p>新規登録(無料)して利用を開始しましょう。</p>
      <div>
        <input className="w-full" type="text" placeholder="ユーザー名" onChange={(e) => setName(e.target.value)} /><br />
        <input className="w-full" type="email" placeholder="メールアドレス" onChange={(e) => setEmail(e.target.value)} /><br />
        <input className="w-full" type="password" placeholder="パスワード" onChange={(e) => setPassword(e.target.value)} /><br />
        <input type="checkbox" />利用規約に同意する<br />
        <input type="checkbox" />プライバシーポリシーに同意する<br />
        <button onClick={() => handleSubmit()} className="bg-green-600 hover:bg-red-500 text-white rounded px-4 py-2">登録する</button>
        <button onClick={() => getUser()} className="bg-blue-600 hover:bg-red-500 text-white rounded px-4 py-2">ユーザー情報の取得</button>
        <button onClick={() => login()} className="bg-blue-600 hover:bg-red-500 text-white rounded px-4 py-2">ログイン</button>
        <button onClick={() => logout()} className="bg-blue-600 hover:bg-red-500 text-white rounded px-4 py-2">ログアウト</button>
        <button onClick={() => mypage()} className="bg-blue-600 hover:bg-red-500 text-white rounded px-4 py-2">認証</button>
      </div>
    </Layout>
  )
}

export default SignUp
