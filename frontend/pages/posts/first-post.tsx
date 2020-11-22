import Head from 'next/head'
import Link from 'next/link'

import Layout from '../../components/layout'

export default function FirstPost() {
    return (
      <Layout>
        <Head>
            <title>Create Next App</title>
        </Head>
        <h1>First Post</h1>
        <h2>
          <Link href="/">
            <a>Back to home</a>
          </Link>
        </h2>
        <img src="/vercel.svg" alt="Vercel Logo" className="logo" />
        <img src="/test.JPG" alt="Vercel Logo" className="logo" />

        <h2>
          外部リンク
            <Link href="/">
            <a className="foo" target="_blank" rel="noopener noreferrer">
                Hello World
            </a>
            </Link>
        </h2>
        <style jsx>{`
            .logo {
                width : 200px;
            }
      `}</style>
      </Layout>

    )


  }