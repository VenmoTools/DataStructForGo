//
// Created by amdins on 19-6-20.
//

#ifndef VIDEO_TREE_H
#define VIDEO_TREE_H

#define TRUE 1
#define FALSE 0

#endif //VIDEO_TREE_H




struct TreeNode;

typedef struct TreeNode *Position;
typedef struct TreeNode *SearchTree;
typedef int ElementType;



/**
 * 将搜索树置空
 * @param T :SearchTree
 * @return NULL
 */
SearchTree Destroy(SearchTree T);

/**
 * 查找元素X
 * @param X 要查找的元素
 * @param T 查找的搜索树
 * @return 查找到的节点
 */
Position Find(ElementType X, SearchTree T);

/**
 *
 * @param T
 * @return
 */
Position FindMax(SearchTree T);

/**
 *
 * @param T
 * @return
 */
Position FindMin(SearchTree T);

/**
 * 在搜索树插入元素，返回插入以后的新树
 * @param T
 * @param X
 * @return
 */
SearchTree Insert(SearchTree T, ElementType X);

SearchTree Delete(SearchTree T, ElementType X);

int GetHeight(SearchTree T);

void PreOrder(SearchTree T);

struct TreeNode {
    ElementType Element;
    SearchTree Left;
    SearchTree Right;
    int height;
};



