import fs from 'fs'
import path from 'path'
import matter from 'gray-matter'
import remark from 'remark'
import html from 'remark-html'

const postDirectory = path.join(process.cwd(),'posts')

export function getSortedPostsData(){
    const fileNames = fs.readdirSync(postDirectory)

    const allPostsData = fileNames.map(fileName => {
        // Remove ".md" from file name to get id
        const id = fileName.replace(/\.md$/, '')

        const fullPath = path.join(postDirectory,fileName)
        const fileContents = fs.readFileSync(fullPath,'utf8')
        const matterResult = matter(fileContents)

        return {
            id,
            ...matterResult.data
        }
    })

    return allPostsData.sort((a,b) => {
        if(a.data < b.data){
            return 1
        } else {
            return -1
        }
    })

}


export function getAllPostIds(){
    const fileNames = fs.readdirSync(postDirectory)

    return fileNames.map(fileName => {
        return {
          params: {
            id: fileName.replace(/\.md$/, '')
          }
        }
    })
}

export async function getPostData(id){
    const fullPath = path.join(postDirectory,`${id}.md`)
    const fileContents = fs.readFileSync(fullPath, 'utf8')

    const matterResult = matter(fileContents)

    const processedContent = await remark()
    .use(html)
    .process(matterResult.content)
    const contentHtml = processedContent.toString()
    return {
        id,
        contentHtml,
        ...matterResult.data
    }
}