import Nav from "./Nav"

const Login = () => {
  return (
    <div>
       <Nav />
       <div className='justify-center flex align-middle pt-10 pb-10 bg-red-200'>
      
      <main>
        <form className='w-100 max-w-96 p-4 m-auto'>
          <h1>Please sign in</h1>
          <input className='relative border  border-black h-auto p-3 font-bold' type='email' id='inputEmail' placeholder='Email address' required />
          <input className='relative border  border-black h-auto p-3 font-bold' type='password' id='inputPassword' placeholder='Password' required />
          <br></br><button className='border border-black w-full m-5' type='submit'>Sign in</button>
        </form>
      </main>
    </div>
    </div>
   
  )
}

export default Login