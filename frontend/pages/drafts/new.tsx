import { useState } from 'react'
import ReactMarkdown from 'react-markdown'
import Header from '../../components/Header'

const New: React.FC = ({ }) => {
  const [text, setText] = useState("")
  return (
    <div>
      <div className='flex flex-col gap-1'>
        <Header showNavBars={false} />
        <input type="text" className='border border-gray-300 border-solid w-[calc(100%-.4rem)] h-12 rounded p-2 text-lg' placeholder='タイトル' />
        <input type="text" className='border border-gray-300 border-solid w-full h-6' placeholder='知識に関連するタグをスペース区切りで5つまで入力（例: Ruby Rails）' />
        <div className='flex w-full h-[calc(36rem)]'>
          <div className="w-2/4 border border-gray-300 border-solid border-collapse">
            <textarea className="w-full h-full" onChange={(e) => setText(e.target.value)}></textarea>
          </div>
          <div className="w-2/4 border border-gray-300 border-solid border-collapse bg-white">
            <ReactMarkdown children={text} />
          </div>
        </div>
        <div className="ml-auto mr-2">
          <div>
            <span className="text-green-400">最終保存: 2022/09/07 17:13:23</span>
            <button className="bg-green-500 hover:bg-green-400 text-white px-4 py-2 ml-2">Qiitaに投稿</button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default New
