import axios from 'axios'
import { useState } from 'react'
import ReactMarkdown from 'react-markdown'
import { useRecoilValue } from 'recoil'
import Header from '../../components/Header'
import { userState } from '../../contexts/UserContext'

const New: React.FC = ({ }) => {
  const [title, setTitle] = useState("")
  const [tags, setTags] = useState<string[]>([])
  const [body, setBody] = useState("")
  const user = useRecoilValue(userState);
  const handleSubmit = async () => {
    try {
      await axios.post("/api/post",
        {
          title,
          body,
          postedBy: user?.id,
          organizationId: user?.organizationId,
          isDraft: 0,
          isDeleted: 0
        })
      tags.forEach(async (tag) => {
        await axios.post("/api/tag", { id: tag, imgUrl: "" })
      })
    } catch (e) {
      console.error(e)
    }
  }
  if (!user) return <></>
  return (
    <div>
      <div className='flex flex-col gap-1'>
        <Header showNavBars={false} />
        <input type="text" className='border border-gray-300 border-solid w-[calc(100%-.4rem)] h-12 rounded p-2 text-lg' placeholder='タイトル' onChange={(e) => setTitle(e.target.value)} />
        <input type="text" className='border border-gray-300 border-solid w-full h-6' placeholder='知識に関連するタグをスペース区切りで5つまで入力（例: Ruby Rails）' onChange={(e) => { setTags(e.target.value.split(" ")) }} />
        <div className='flex w-full h-[calc(36rem)]'>
          <div className="w-2/4 border border-gray-300 border-solid border-collapse">
            <textarea placeholder="エンジニアに関わる知識をMarkdown記法で書いて共有しよう" className="w-full h-full" onChange={(e) => setBody(e.target.value)}></textarea>
          </div>
          <div className="w-2/4 border border-gray-300 border-solid border-collapse bg-white">
            <ReactMarkdown children={body} />
          </div>
        </div>
        <div className="ml-auto mr-2">
          <div>
            <button onClick={handleSubmit} className="bg-green-500 hover:bg-green-400 text-white px-4 py-2 ml-2">Qiitaに投稿</button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default New
