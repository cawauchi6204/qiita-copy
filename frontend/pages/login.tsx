import Layout from "../components/Layout"
import { useForm, SubmitHandler } from "react-hook-form";
import axios from "axios";
import Link from "next/link";

type Inputs = {
  email: string
  password: string
};

type Props = {}

const Login: React.FC<Props> = ({ }) => {
  const {
    register,
    handleSubmit,
  } = useForm<Inputs>();
  const onSubmit: SubmitHandler<Inputs> = async (data) => {
    await axios.post('/api/login', data)
  };
  return (
    <Layout>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="mt-8">
          <input type="email" {...register("email")} className="w-full border border-solid border-gray p-2" placeholder="メールアドレス" />
        </div>
        <div className="mt-8">
          <input type="password" {...register("password")} className="w-full border border-solid border-gray p-2" placeholder="パスワード" />
        </div>
        <button type="submit" className="cursor-pointer py-2 px-4 mt-2 rounded text-white" style={{ "background": "#55C500" }}>qiitaにログイン</button>
      </form>
      <p className="text-xs mt-4">Qiita のアカウントを持っていない場合は<Link href="/signup"><a className="text-blue-500">新規登録</a></Link>から。</p>
    </Layout>
  )
}

export default Login
