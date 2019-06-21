//
// Created by amdins on 19-6-20.
//

#include "tree.h"
#include <stdio.h>
#include <stdlib.h>


SearchTree Destroy(SearchTree T) {
    if (T != NULL) {
        Destroy(T->Left);
        Destroy(T->Right);
        free(T);
    }
    return NULL;
}

Position Find(ElementType X, SearchTree T) {
    if (T == NULL)
        return NULL;
    if (X < T->Element)
        return Find(X, T->Left);
    else if (X > T->Element)
        return Find(X, T->Right);
    else
        return T;
}

int GetHeight(SearchTree T) {
    if (T == NULL)
        return 0;
    return T->height;
}

Position FindMax(SearchTree T) {
    if (T == NULL)
        return NULL;
    if (T->Right == NULL)
        return T;
    return FindMax(T->Right);
}

Position FindMin(SearchTree T) {
    if (T == NULL)
        return NULL;
    if (T->Left == NULL)
        return T;
    return FindMin(T->Left);
}

int Max(int a, int b) {
    if (a > b)
        return a;
    return b;
}

/**
 * 计算平衡因子,当前节点左子树的高度 - 右子树的高度
 * @param T 当前节点
 * @return 平衡因子
 */
int getBalanceFactor(SearchTree T) {
    if (T == NULL)
        return 0;
    return GetHeight(T->Left) - GetHeight(T->Right);
}

_Bool isSearchTree(SearchTree T) {


    return TRUE;
}

void inOrder(SearchTree T, int *arr) {
    if (arr == NULL)
        return;

    inOrder(T->Left, arr);

}

SearchTree RightRotate(SearchTree T) {

    /*
     *  Before
     *          T
     *        /   \
     *       T2    z
     *      /  \
     *     y    x
     *    /
     *   g
     */

    SearchTree T2 = T->Left;
    SearchTree x = T2->Right;
    T2->Right = T;
    T->Left = x;

    /*
     * After
     *       T2
     *      /  \
     *     y    T
     *    /    / \
     *   g    x   z
     */
    T->height = Max(GetHeight(T->Left), GetHeight(T->Right)) + 1;
    T2->height = Max(GetHeight(T2->Left), GetHeight(T->Right)) + 1;

    return T2;
}

SearchTree LeftRotate(SearchTree T) {
    /*
     * Before
     *        T
     *      /   \
     *     z    T2
     *         /  \
     *        x    y
     *              \
     *               g
     */
    SearchTree T2 = T->Right;
    SearchTree x = T2->Left;
    T2->Left = T;
    T->Right = x;

    /*
     * After
     *
     *          T2
     *         /  \
     *        T    y
     *       / \    \
     *      z   x    g
     */
    T->height = Max(GetHeight(T->Left), GetHeight(T->Right)) + 1;
    T2->height = Max(GetHeight(T->Left), GetHeight(T->Right)) + 1;

    return T2;
}


SearchTree Insert(SearchTree T, ElementType X) {
    if (T == NULL) {
        T = malloc(sizeof(struct TreeNode));
        if (T == NULL) {
            fprintf(stderr, "Out of space!");
            exit(EXIT_FAILURE);
        }
        T->Left = T->Right = NULL;
        T->Element = X;
        T->height = 1;
    } else {
        if (X < T->Element)
            T->Left = Insert(T->Left, X);
        else
            T->Right = Insert(T->Right, X);
    }
    // 更新height值 当前节点height值为左子树节点和右子树节点中最大的一个+1
    T->height = 1 + Max(GetHeight(T->Left), GetHeight(T->Right));

    // 获取当前节点的平衡因子
    int balanceFactor = getBalanceFactor(T);
    if (abs(balanceFactor) > 1)
        fprintf(stdout, "current:%d balance: %d\n", T->Element, balanceFactor);

    /*
     * Left
     *     h:节点高度
     *     b:平衡因子
     *             T (h=4 b=2)
     *           /   \
     *(h=3 b=1) T2    z (h=1)
     *         /  \
     * (h=2)  y    x (h=1)
     *       / \
     * (h=1)g   h
     */
    if (balanceFactor > 1 && getBalanceFactor(T->Left) >= 0) {
        return RightRotate(T);
    }
    /*
     * Right
     *        T (h=4 b=-2) 1-3 = -2
     *      /   \
     *(h=1)z    T2 (h=3 b=-1) -> 1-2
     *         /  \
     *  (h=1) x    y (h=2 b=-1)
     *            / \
     *           h   g (h=1)
     */
    if (balanceFactor < -1 && getBalanceFactor(T->Right) <= 0) {
        return LeftRotate(T);
    }

    /*
     *  (h=5 b=3)     T
     *              /    \
     *  (h=4 b=2) T3     T2
     *           /  \    / \
     *(h=3 b=1) T4   z  x   y
     *         /  \
     * (h=2) T5   a (h=1)
     *         \
     *          g (h=1)
     */
    // LR
    if (balanceFactor > 1 && getBalanceFactor(T->Left) < 0) {
        T->Left = LeftRotate(T->Left);
        return RightRotate(T);
    }
    //RL
    if (balanceFactor < 1 && getBalanceFactor(T->Right) > 0) {
        T->Right = RightRotate(T->Right);
        return LeftRotate(T);
    }
    return T;
}


SearchTree Delete(SearchTree T, ElementType X) {
    Position TempCell;

    if (T == NULL) {
        fprintf(stderr, "Element Not Found");
        exit(EXIT_FAILURE);
    }
    if (X < T->Element)
        T->Left = Delete(T->Left, X);
    else
        T->Right = Delete(T->Right, X);
    // 该节点有左孩子并且有右孩子
    if (T->Right && T->Left) {
        //使用tempCell获取右子树的最小元素
        TempCell = FindMin(T->Right);
        // 替换该孩子的节点值
        T->Element = TempCell->Element;
        T->Right = Delete(T->Right, X);
    } else {
        // 只有一个孩子或没有孩子
        TempCell = T;
        if (T->Left == NULL)
            T = T->Right;
        else if (T->Right == NULL)
            T = T->Left;
        free(TempCell);
    }

    return T;
}

void PreOrder(SearchTree T) {

    if (T == NULL)
        return;

    printf("%d\n", T->Element);
    PreOrder(T->Left);
    PreOrder(T->Right);
}
