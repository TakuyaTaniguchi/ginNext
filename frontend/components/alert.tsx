
import cn from 'classnames'
import test from '../styles/test.module.scss'

export default function Alert({ children, type }) {
  return (
    <div
      className={cn({
        [`success`]: type === 'success',
        [`error`]: type === 'error'
      })}
    >
        <p className={test.test}>test</p>
        <p className={test.test_hoge}>Hoge</p>

      {children}
      <style jsx>{`
      .success {
            color: green;
        }
        .error {
            color: red;
        }
      `}</style>
    </div>
  )

}