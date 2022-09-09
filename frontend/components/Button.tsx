type Props = {
  text: string
  color: string
}

const Button: React.FC<Props> = ({ text, color }) => {
  return <button type="button" className={`inline-block px-6 py-2.5 bg-${color}-500 text-white text-xs font-bold leading-tight uppercase hover:bg-${color}-600  focus:bg-${color}-900 w-full focus:shadow-lg focus:outline-none focus:ring-0 active:bg-${color}-900 active:shadow-lg transition duration-150 ease-in-out`}>{text}</button>
}

export default Button
